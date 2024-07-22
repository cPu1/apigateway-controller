	if delta.DifferentAt("Spec.Tags") {
		if err := syncTags(
			ctx, rm.sdkapi, rm.metrics,
			makeARN(*desired.ko.Status.ID),
			desired.ko.Spec.Tags, latest.ko.Spec.Tags,
		); err != nil {
			return nil, err
		}
	} else if !delta.DifferentExcept("Spec.Tags") {
        return desired, nil
    }
