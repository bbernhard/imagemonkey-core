package tests

import (
	"testing"
	"os/exec"	
	"time"
	"os"
	"strconv"
	"github.com/bbernhard/imagemonkey-core/commons"
)

func runLabelsDownloader(t *testing.T) {
	redisAddress := ":6379"
	if commons.GetEnv("REDIS_ADDRESS") != "" {
		redisAddress = commons.GetEnv("REDIS_ADDRESS")
	}

	os.RemoveAll("/tmp/labels-unittest-backups")
	// Start a process
	cmd := exec.Command("go", "run", "-tags", "dev", "labels_downloader.go", "-autoclose_github_issue=false", 
						"-singleshot=true", "-labels_dir=/tmp/labels-unittest", "-backup_dir=/tmp/labels-unittest-backups", 
						"-labels_repository_url=/tmp/labels-unittest", "-use_backup_timestamp=false",
						"-download_dir=/tmp/labels-unittest-backups", "-redis_address="+redisAddress)
	cmd.Dir = "../src"
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	ok(t, err)

	// Wait for the process to finish or kill it after a timeout:
	done := make(chan error, 1)
	go func() {
	    done <- cmd.Wait()
	}()
	select {
	case <-time.After(60 * time.Second):
	    err := cmd.Process.Kill()
	    ok(t, err) //failed to kill process
	    t.Errorf("process killed as timeout reached")
	case err := <-done:
	    ok(t, err)
	}
}

func verifyProductiveAnnotations(t *testing.T, imageAnnotationSuggestionEntries []ImageAnnotationEntry, imageAnnotationEntries []ImageAnnotationEntry, 
									annotationSuggestionDataEntries []AnnotationDataEntry, annotationDataEntries []AnnotationDataEntry,
									imageAnnotationSuggestionRevisionEntries []ImageAnnotationRevisionEntry, 
									imageAnnotationRevisionEntries []ImageAnnotationRevisionEntry) {
	equals(t, len(imageAnnotationSuggestionEntries), len(imageAnnotationEntries))
	equals(t, len(annotationSuggestionDataEntries), len(annotationDataEntries))
	equals(t, len(imageAnnotationSuggestionRevisionEntries), len(imageAnnotationRevisionEntries))

	imageAnnotationIdRevisionMapping := make(map[int64]int64)
	imageAnnotationSuggestionIdRevisionMapping := make(map[int64]int64)
	imageAnnotationRevisionIdMapping := make(map[int64]int64)
	for i, imageAnnotationRevisionEntry := range imageAnnotationRevisionEntries {
		equals(t, imageAnnotationRevisionEntry.Revision, imageAnnotationSuggestionRevisionEntries[i].Revision)
		imageAnnotationIdRevisionMapping[imageAnnotationRevisionEntry.Id] = imageAnnotationRevisionEntry.ImageAnnotationId
		imageAnnotationSuggestionIdRevisionMapping[imageAnnotationSuggestionRevisionEntries[i].Id] = imageAnnotationSuggestionRevisionEntries[i].ImageAnnotationId
		imageAnnotationRevisionIdMapping[imageAnnotationSuggestionRevisionEntries[i].Id] = imageAnnotationRevisionEntry.Id
	}

	imageAnnotationIdMapping := make(map[int64]int64)
	for i, imageAnnotationEntry := range imageAnnotationEntries {
		equals(t, imageAnnotationEntry.Uuid, imageAnnotationSuggestionEntries[i].Uuid)
		equals(t, imageAnnotationEntry.ImageId, imageAnnotationSuggestionEntries[i].ImageId)
		equals(t, imageAnnotationEntry.NumOfValid, imageAnnotationSuggestionEntries[i].NumOfValid)
		equals(t, imageAnnotationEntry.NumOfInvalid, imageAnnotationSuggestionEntries[i].NumOfInvalid)
		equals(t, imageAnnotationEntry.FingerprintOfLastModification, imageAnnotationSuggestionEntries[i].FingerprintOfLastModification)
		equals(t, imageAnnotationEntry.AutoGenerated, imageAnnotationSuggestionEntries[i].AutoGenerated)
		equals(t, imageAnnotationEntry.Revision, imageAnnotationSuggestionEntries[i].Revision)

		imageAnnotationIdMapping[imageAnnotationSuggestionEntries[i].Id] = imageAnnotationEntry.Id

		labelName, err := db.GetLabelSuggestionNameFromId(imageAnnotationSuggestionEntries[i].LabelId)
		ok(t, err)
		labelId, err := db.GetLabelIdFromName(labelName)
		ok(t, err)
		equals(t, imageAnnotationEntry.LabelId, labelId)
	}

	for i, annotationDataEntry := range annotationDataEntries {
		equals(t, annotationDataEntry.Uuid, annotationSuggestionDataEntries[i].Uuid)
		
		if annotationDataEntry.ImageAnnotationId == -1 {
			equals(t, int64(annotationSuggestionDataEntries[i].ImageAnnotationId), int64(-1))

			imgAnnotationId := imageAnnotationIdRevisionMapping[annotationDataEntry.ImageAnnotationRevisionId]
			imgAnnotationSuggestionId := imageAnnotationSuggestionIdRevisionMapping[annotationSuggestionDataEntries[i].ImageAnnotationRevisionId]
			equals(t, imgAnnotationId, imageAnnotationIdMapping[imgAnnotationSuggestionId])
		} else {
			equals(t, annotationDataEntry.ImageAnnotationId, imageAnnotationIdMapping[annotationSuggestionDataEntries[i].ImageAnnotationId])
		}

		equals(t, annotationDataEntry.AnnotationTypeId, annotationSuggestionDataEntries[i].AnnotationTypeId)
		equals(t, annotationDataEntry.ImageAnnotationRevisionId, imageAnnotationRevisionIdMapping[annotationSuggestionDataEntries[i].ImageAnnotationRevisionId])
		equals(t, annotationDataEntry.Annotation, annotationSuggestionDataEntries[i].Annotation)
	}
}

func TestLabelsDownloaderSuccess(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	testSignUp(t, "testuser", "testpassword", "testuser@imagemonkey.io")
	token := testLogin(t, "testuser", "testpassword", 200)

	err := db.GiveUserModeratorRights("testuser")
	ok(t, err)

	testMultipleDonate(t, "floor")

	imageIds, err := db.GetAllImageIds()
	ok(t, err)

	for _, imageId := range imageIds {
		testSuggestLabelForImage(t, imageId, "red apple", true, token, 200)
	}
	runTrendingLabelsWorker(t, 5)

	trendingLabels := testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 1)


	testAcceptTrendingLabel(t, "red apple", "", "red apples", "red apple", "", token, "normal", 201)
	trendingLabels = testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 1)

	equals(t, trendingLabels[0].Status, "accepted")

	runLabelBot(t, "cisuccess")
	trendingLabels = testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 1)
	equals(t, trendingLabels[0].Status, "merged")

	numberOfLabelsBefore, err := db.GetNumberOfLabels()
	ok(t, err)
	runLabelsDownloader(t)
	numberOfLabelsAfter, err := db.GetNumberOfLabels()
	ok(t, err)
	equals(t, numberOfLabelsBefore+1, numberOfLabelsAfter)
}

func TestLabelsDownloaderFailure(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	testSignUp(t, "testuser", "testpassword", "testuser@imagemonkey.io")
	token := testLogin(t, "testuser", "testpassword", 200)

	err := db.GiveUserModeratorRights("testuser")
	ok(t, err)

	testMultipleDonate(t, "floor")

	imageIds, err := db.GetAllImageIds()
	ok(t, err)

	for _, imageId := range imageIds {
		testSuggestLabelForImage(t, imageId, "red apple", true, token, 200)
	}
	runTrendingLabelsWorker(t, 5)

	trendingLabels := testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 1)


	testAcceptTrendingLabel(t, "red apple", "", "red apples", "red apple", "", token, "normal", 201)
	trendingLabels = testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 1)

	equals(t, trendingLabels[0].Status, "accepted")

	runLabelBot(t, "cifailure")
	trendingLabels = testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 1)
	equals(t, trendingLabels[0].Status, "build-failed")

	numberOfLabelsBefore, err := db.GetNumberOfLabels()
	ok(t, err)
	runLabelsDownloader(t)
	numberOfLabelsAfter, err := db.GetNumberOfLabels()
	ok(t, err)
	equals(t, numberOfLabelsBefore, numberOfLabelsAfter)
}

func TestLabelsDownloaderFailureCanBeRetried(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	testSignUp(t, "testuser", "testpassword", "testuser@imagemonkey.io")
	token := testLogin(t, "testuser", "testpassword", 200)

	err := db.GiveUserModeratorRights("testuser")
	ok(t, err)

	testMultipleDonate(t, "floor")

	imageIds, err := db.GetAllImageIds()
	ok(t, err)

	for _, imageId := range imageIds {
		testSuggestLabelForImage(t, imageId, "red apple", true, token, 200)
	}
	runTrendingLabelsWorker(t, 5)

	trendingLabels := testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 1)


	testAcceptTrendingLabel(t, "red apple", "", "red apples", "red apple", "", token, "normal", 201)
	trendingLabels = testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 1)

	equals(t, trendingLabels[0].Status, "accepted")

	runLabelBot(t, "cifailure")
	trendingLabels = testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 1)
	equals(t, trendingLabels[0].Status, "build-failed")

	numberOfLabelsBefore, err := db.GetNumberOfLabels()
	ok(t, err)
	runLabelsDownloader(t)
	numberOfLabelsAfter, err := db.GetNumberOfLabels()
	ok(t, err)
	equals(t, numberOfLabelsBefore, numberOfLabelsAfter)

	testAcceptTrendingLabel(t, "red apple", "", "red apples", "red apple", "", token, "normal", 201) 
	trendingLabels = testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 1)

	equals(t, trendingLabels[0].Status, "retry")
}

func TestLabelsDownloaderSuccessCannotBeRetried(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	testSignUp(t, "testuser", "testpassword", "testuser@imagemonkey.io")
	token := testLogin(t, "testuser", "testpassword", 200)

	err := db.GiveUserModeratorRights("testuser")
	ok(t, err)

	testMultipleDonate(t, "floor")

	imageIds, err := db.GetAllImageIds()
	ok(t, err)

	for _, imageId := range imageIds {
		testSuggestLabelForImage(t, imageId, "red apple", true, token, 200)
	}
	runTrendingLabelsWorker(t, 5)

	trendingLabels := testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 1)


	testAcceptTrendingLabel(t, "red apple", "", "red apples", "red apple", "", token, "normal", 201)
	trendingLabels = testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 1)

	equals(t, trendingLabels[0].Status, "accepted")

	runLabelBot(t, "cisuccess")
	trendingLabels = testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 1)
	equals(t, trendingLabels[0].Status, "merged")

	numberOfLabelsBefore, err := db.GetNumberOfLabels()
	ok(t, err)
	runLabelsDownloader(t)
	numberOfLabelsAfter, err := db.GetNumberOfLabels()
	ok(t, err)
	equals(t, numberOfLabelsBefore+1, numberOfLabelsAfter)

	testAcceptTrendingLabel(t, "red apple", "", "red apples", "red apple", "", token, "normal", 201) 
	trendingLabels = testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 1)

	equals(t, trendingLabels[0].Status, "productive")
}


func TestLabelsDownloaderNonProductiveLabelWithAnnotationsSuccess(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	testSignUp(t, "testuser", "testpassword", "testuser@imagemonkey.io")
	token := testLogin(t, "testuser", "testpassword", 200)

	err := db.GiveUserModeratorRights("testuser")
	ok(t, err)

	testMultipleDonate(t, "floor")

	imageIds, err := db.GetAllImageIds()
	ok(t, err)

	for _, imageId := range imageIds {
		testSuggestLabelForImage(t, imageId, "red apple", true, token, 200)
	}

	//annotate label suggestions
	for _, imageId := range imageIds {
		testAnnotate(t, imageId, "red apple", "", 
					`[{"top":50,"left":300,"type":"rect","angle":15,"width":240,"height":100,"stroke":{"color":"red","width":1}}]`, token, 201)
	}

	imageAnnotationSuggestionEntries, err := db.GetImageAnnotationSuggestionEntries()
	ok(t, err)

	annotationSuggestionDataEntries, err := db.GetAnnotationSuggestionDataEntries()
	ok(t, err)

	imageAnnotationSuggestionRevisionEntries, err := db.GetImageAnnotationSuggestionRevisionEntries()
	ok(t, err)

	runTrendingLabelsWorker(t, 5)

	trendingLabels := testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 1)


	testAcceptTrendingLabel(t, "red apple", "", "red apples", "red apple", "", token, "normal", 201)
	trendingLabels = testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 1)

	equals(t, trendingLabels[0].Status, "accepted")

	runLabelBot(t, "cisuccess")
	trendingLabels = testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 1)
	equals(t, trendingLabels[0].Status, "merged")

	numberOfLabelsBefore, err := db.GetNumberOfLabels()
	ok(t, err)
	runLabelsDownloader(t)
	numberOfLabelsAfter, err := db.GetNumberOfLabels()
	ok(t, err)
	equals(t, numberOfLabelsBefore+1, numberOfLabelsAfter)

	imageAnnotationEntries, err := db.GetImageAnnotationEntries()
	ok(t, err)

	annotationDataEntries, err := db.GetAnnotationDataEntries()
	ok(t, err)

	imageAnnotationRevisionEntries, err := db.GetImageAnnotationRevisionEntries()
	ok(t, err)

	verifyProductiveAnnotations(t, imageAnnotationSuggestionEntries, imageAnnotationEntries, annotationSuggestionDataEntries, annotationDataEntries,
								imageAnnotationSuggestionRevisionEntries, imageAnnotationRevisionEntries)
}


func TestLabelsDownloaderMultipleNonProductiveLabelWithAnnotationsSuccess(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	testSignUp(t, "testuser", "testpassword", "testuser@imagemonkey.io")
	token := testLogin(t, "testuser", "testpassword", 200)

	err := db.GiveUserModeratorRights("testuser")
	ok(t, err)

	testMultipleDonate(t, "floor")

	imageIds, err := db.GetAllImageIds()
	ok(t, err)

	for i, imageId := range imageIds {
		testSuggestLabelForImage(t, imageId, "red apple " + strconv.Itoa(i), true, token, 200)
	}

	//annotate label suggestions
	for i, imageId := range imageIds {
		testAnnotate(t, imageId, "red apple " + strconv.Itoa(i), "", 
					`[{"top":50,"left":300,"type":"rect","angle":15,"width":240,"height":100,"stroke":{"color":"red","width":1}}]`, token, 201)
	}

	imageAnnotationSuggestionEntries, err := db.GetImageAnnotationSuggestionEntries()
	ok(t, err)

	annotationSuggestionDataEntries, err := db.GetAnnotationSuggestionDataEntries()
	ok(t, err)

	imageAnnotationSuggestionRevisionEntries, err := db.GetImageAnnotationSuggestionRevisionEntries()
	ok(t, err)

	runTrendingLabelsWorker(t, 0)

	trendingLabels := testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 13)

	for i, _ := range imageIds {
		testAcceptTrendingLabel(t, "red apple " + strconv.Itoa(i), "", "red apples", "red apple " + strconv.Itoa(i), "", token, "normal", 201)
	}

	trendingLabels = testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 13)

	equals(t, trendingLabels[0].Status, "accepted")

	runLabelBot(t, "cisuccess")
	trendingLabels = testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 13)
	equals(t, trendingLabels[0].Status, "merged")

	numberOfLabelsBefore, err := db.GetNumberOfLabels()
	ok(t, err)
	runLabelsDownloader(t)
	numberOfLabelsAfter, err := db.GetNumberOfLabels()
	ok(t, err)
	equals(t, numberOfLabelsBefore+13, numberOfLabelsAfter)

	imageAnnotationEntries, err := db.GetImageAnnotationEntries()
	ok(t, err)

	annotationDataEntries, err := db.GetAnnotationDataEntries()
	ok(t, err)

	imageAnnotationRevisionEntries, err := db.GetImageAnnotationRevisionEntries()
	ok(t, err)

	verifyProductiveAnnotations(t, imageAnnotationSuggestionEntries, imageAnnotationEntries, annotationSuggestionDataEntries, annotationDataEntries,
								imageAnnotationSuggestionRevisionEntries, imageAnnotationRevisionEntries)
}


func TestLabelsDownloaderMultipleNonProductiveLabelWithAnnotationsAndRefinementsSuccess(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	testSignUp(t, "testuser", "testpassword", "testuser@imagemonkey.io")
	token := testLogin(t, "testuser", "testpassword", 200)

	err := db.GiveUserModeratorRights("testuser")
	ok(t, err)

	testMultipleDonate(t, "floor")

	imageIds, err := db.GetAllImageIds()
	ok(t, err)

	for i, imageId := range imageIds {
		testSuggestLabelForImage(t, imageId, "red apple " + strconv.Itoa(i), true, token, 200)
	}

	//annotate label suggestions
	for i, imageId := range imageIds {
		testAnnotate(t, imageId, "red apple " + strconv.Itoa(i), "", 
					`[{"top":50,"left":300,"type":"rect","angle":15,"width":240,"height":100,"stroke":{"color":"red","width":1}}]`, token, 201)
		annotationIds, err := db.GetImageAnnotationSuggestionIdsForImage(imageId)
		ok(t, err)
		equals(t, len(annotationIds), 1)
		newAnnotations := `[{"top":55,"left":310,"type":"rect","angle":16,"width":244,"height":120,"stroke":{"color":"red","width":1}}]`
		testAnnotationRework(t, annotationIds[0], newAnnotations, token)
	}

	imageAnnotationSuggestionEntries, err := db.GetImageAnnotationSuggestionEntries()
	ok(t, err)

	annotationSuggestionDataEntries, err := db.GetAnnotationSuggestionDataEntries()
	ok(t, err)

	imageAnnotationSuggestionRevisionEntries, err := db.GetImageAnnotationSuggestionRevisionEntries()
	ok(t, err)

	runTrendingLabelsWorker(t, 0)

	trendingLabels := testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 13)

	for i, _ := range imageIds {
		testAcceptTrendingLabel(t, "red apple " + strconv.Itoa(i), "", "red apples", "red apple " + strconv.Itoa(i), "", token, "normal", 201)
	}

	trendingLabels = testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 13)

	equals(t, trendingLabels[0].Status, "accepted")

	runLabelBot(t, "cisuccess")
	trendingLabels = testGetTrendingLabels(t, token, 200)
	equals(t, len(trendingLabels), 13)
	equals(t, trendingLabels[0].Status, "merged")

	numberOfLabelsBefore, err := db.GetNumberOfLabels()
	ok(t, err)
	runLabelsDownloader(t)
	numberOfLabelsAfter, err := db.GetNumberOfLabels()
	ok(t, err)
	equals(t, numberOfLabelsBefore+13, numberOfLabelsAfter)

	imageAnnotationEntries, err := db.GetImageAnnotationEntries()
	ok(t, err)

	annotationDataEntries, err := db.GetAnnotationDataEntries()
	ok(t, err)

	imageAnnotationRevisionEntries, err := db.GetImageAnnotationRevisionEntries()
	ok(t, err)

	verifyProductiveAnnotations(t, imageAnnotationSuggestionEntries, imageAnnotationEntries, annotationSuggestionDataEntries, annotationDataEntries,
								imageAnnotationSuggestionRevisionEntries, imageAnnotationRevisionEntries)
}
