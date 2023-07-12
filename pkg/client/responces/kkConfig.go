package responces

import "encoding/json"

func UnmarshalKkConfig(data []byte) (KkConfig, error) {
	var r KkConfig
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *KkConfig) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type KkConfig struct {
	Issuer                                             string   `json:"issuer"`
	AuthorizationEndpoint                              string   `json:"authorization_endpoint"`
	TokenEndpoint                                      string   `json:"token_endpoint"`
	IntrospectionEndpoint                              string   `json:"introspection_endpoint"`
	UserinfoEndpoint                                   string   `json:"userinfo_endpoint"`
	EndSessionEndpoint                                 string   `json:"end_session_endpoint"`
	JwksURI                                            string   `json:"jwks_uri"`
	CheckSessionIframe                                 string   `json:"check_session_iframe"`
	GrantTypesSupported                                []string `json:"grant_types_supported"`
	ResponseTypesSupported                             []string `json:"response_types_supported"`
	SubjectTypesSupported                              []string `json:"subject_types_supported"`
	IDTokenSigningAlgValuesSupported                   []string `json:"id_token_signing_alg_values_supported"`
	IDTokenEncryptionAlgValuesSupported                []string `json:"id_token_encryption_alg_values_supported"`
	IDTokenEncryptionEncValuesSupported                []string `json:"id_token_encryption_enc_values_supported"`
	UserinfoSigningAlgValuesSupported                  []string `json:"userinfo_signing_alg_values_supported"`
	RequestObjectSigningAlgValuesSupported             []string `json:"request_object_signing_alg_values_supported"`
	ResponseModesSupported                             []string `json:"response_modes_supported"`
	RegistrationEndpoint                               string   `json:"registration_endpoint"`
	TokenEndpointAuthMethodsSupported                  []string `json:"token_endpoint_auth_methods_supported"`
	TokenEndpointAuthSigningAlgValuesSupported         []string `json:"token_endpoint_auth_signing_alg_values_supported"`
	IntrospectionEndpointAuthMethodsSupported          []string `json:"introspection_endpoint_auth_methods_supported"`
	IntrospectionEndpointAuthSigningAlgValuesSupported []string `json:"introspection_endpoint_auth_signing_alg_values_supported"`
	ClaimsSupported                                    []string `json:"claims_supported"`
	ClaimTypesSupported                                []string `json:"claim_types_supported"`
	ClaimsParameterSupported                           bool     `json:"claims_parameter_supported"`
	ScopesSupported                                    []string `json:"scopes_supported"`
	RequestParameterSupported                          bool     `json:"request_parameter_supported"`
	RequestURIParameterSupported                       bool     `json:"request_uri_parameter_supported"`
	RequireRequestURIRegistration                      bool     `json:"require_request_uri_registration"`
	CodeChallengeMethodsSupported                      []string `json:"code_challenge_methods_supported"`
	TLSClientCertificateBoundAccessTokens              bool     `json:"tls_client_certificate_bound_access_tokens"`
	RevocationEndpoint                                 string   `json:"revocation_endpoint"`
	RevocationEndpointAuthMethodsSupported             []string `json:"revocation_endpoint_auth_methods_supported"`
	RevocationEndpointAuthSigningAlgValuesSupported    []string `json:"revocation_endpoint_auth_signing_alg_values_supported"`
	BackchannelLogoutSupported                         bool     `json:"backchannel_logout_supported"`
	BackchannelLogoutSessionSupported                  bool     `json:"backchannel_logout_session_supported"`
	DeviceAuthorizationEndpoint                        string   `json:"device_authorization_endpoint"`
	BackchannelTokenDeliveryModesSupported             []string `json:"backchannel_token_delivery_modes_supported"`
	BackchannelAuthenticationEndpoint                  string   `json:"backchannel_authentication_endpoint"`
}
