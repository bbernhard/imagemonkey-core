function buildComposedLabels(label, uuid, sublabels) {
    if (sublabels === null || sublabels === undefined)
        return [{
            "displayname": label,
            "uuid": uuid
        }];
    composedLabels = [{
        "displayname": label,
        "uuid": uuid
    }];
    for (const sublabel of sublabels) {
        composedLabels.push({
            "displayname": sublabel.name + "/" + label,
            "uuid": sublabel.uuid
        });
    }
    return composedLabels;
}

function labelExistsInLabelList(label, labelList) {
    for (const elem of labelList) {
        if (elem["displayname"] === label)
            return true;
    }
    return false;
}

function removeLabelFromLabelList(label, labelList) {
    let i = 0;
    for (const elem of labelList) {
        if (elem["displayname"] === label) {
            labelList.splice(i, 1);
            break;
        }
        i++;
    }
}