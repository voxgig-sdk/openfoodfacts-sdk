package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewProductEntityFunc func(client *OpenfoodfactsSDK, entopts map[string]any) OpenfoodfactsEntity

var NewSearchEntityFunc func(client *OpenfoodfactsSDK, entopts map[string]any) OpenfoodfactsEntity

