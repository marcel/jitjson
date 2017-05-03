package thirdparty

type ThirdPartyStruct struct {
	// This has a JSON tag but it's in the vendor dir so ignored
	SomeField string `json:"some_field"`
}
