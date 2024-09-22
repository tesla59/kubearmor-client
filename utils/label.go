package utils

import "strings"

// LabelMap is an alias for map[string]string
type LabelMap = map[string]string

func labelSplitter(r rune) bool {
	return r == ':' || r == '='
}

func LabelArrayToLabelMap(labels []string) LabelMap {
	labelMap := LabelMap{}
	for _, label := range labels {
		kvPair := strings.FieldsFunc(label, labelSplitter)
		if len(kvPair) != 2 {
			continue
		}
		labelMap[kvPair[0]] = kvPair[1]
	}
	return labelMap
}
