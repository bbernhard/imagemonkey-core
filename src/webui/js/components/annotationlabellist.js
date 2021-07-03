AnnotationLabelListComponent = {
    template: "#annotation-label-list-template",
    delimiters: ['${', '}$'],
    data() {
        return {
            labels: [],
            labelLookupTable: [],
            currentSelectedItem: null,
            visible: true
        }
    },
    computed: {},
    methods: {
        itemSelected: function(labelUuid) {
            this.currentSelectedItem = labelUuid;
        },
        itemColor: function(labelUuid) {
            if (this.currentSelectedItem === labelUuid)
                return "bg-red-100";
            return "bg-green-100";
        },
        removeLabel: function(label) {
            EventBus.$emit("removeLabel", label);
        },
        getLabelsForImage: function(imageId) {
            var that = this;
            let onlyUnlockedLabels = false;
            imageMonkeyApi.getLabelsForImage(imageId, onlyUnlockedLabels)
                .then(function(entries) {
                    let composedLabels = []
                    for (const entry of entries) {
                        composedLabels.push(...buildComposedLabels(entry.label, entry.uuid, entry.sublabels));
                    }

                    that.labelLookupTable = {}
                    for (const composedLabel of composedLabels) {
                        that.labelLookupTable[composedLabel.uuid] = composedLabel.displayname;
                    }

                    that.labels = composedLabels;
                }).catch(function(e) {
                    console.log(e.message);
                    Sentry.captureException(e);
                });
        },
        getAnnotationsForImage: function(imageId, imageUnlocked) {
            var that = this;
            imageMonkeyApi.getAnnotationsForImage(imageId, imageUnlocked)
                .then(function(annotations) {
                    //TODO
                }).catch(function(e) {
                    console.log(e.message);
                    Sentry.captureException(e);
                });
        },
        onUnannotatedImageDataReceived: function(data) {
            this.getLabelsForImage(data.uuid);
            this.getAnnotationsForImage(data.uuid, data.unlocked);
        }
    },
    beforeDestroy: function() {
        EventBus.$off("unannotatedImageDataReceived", this.onUnannotatedImageDataReceived);
    },
    mounted: function() {
        EventBus.$on("unannotatedImageDataReceived", this.onUnannotatedImageDataReceived);
    }

};
