package gofeatures

//goland:noinspection GoUnusedExportedFunction
func ConvertFeatures(names []string) map[string]bool {
	features := map[string]bool{}
	for _, feature := range names {
		features[feature] = true
	}
	return features
}

func HasFeature(features map[string]bool, name string) bool {
	_, found := features[name]
	return found
}

//goland:noinspection GoUnusedExportedFunction
func HasFeatureAndNoError(features map[string]bool, name string, err error) bool {
	return nil == err && HasFeature(features, name)
}
