package orgchecker

// OrdererConfig ...
type OrdererConfig struct {
	BatchSizeMaxMessageCount   int     `validate:"numeric"`
	BatchSizeAbsoluteMaxBytes  int     `validate:"numeric"`
	BatchSizePrefferedMaxBytes float64 `validate:"numeric"`
	BatchTimeout               string
}

// OrganizationRemove структура содержащая в себе только
type OrganizationRemove struct {
	OrgName string `validate:"required,alphanum"`
}

// OrganizationConfig структура которая содержит в себе информацию об организации и валидирует входящие параметры
type OrganizationConfig struct {
	Groups struct {
	} `json:"groups"`
	ModPolicy string `json:"mod_policy"`
	Policies  struct {
		Admins struct {
			ModPolicy string `json:"mod_policy"`
			Policy    struct {
				Type  int `json:"type"`
				Value struct {
					Identities []struct {
						Principal struct {
							MspIdentifier string `json:"msp_identifier"`
							Role          string `json:"role"`
						} `json:"principal"`
						PrincipalClassification string `json:"principal_classification"`
					} `json:"identities"`
					Rule struct {
						NOutOf struct {
							N     int `json:"n"`
							Rules []struct {
								SignedBy int `json:"signed_by"`
							} `json:"rules"`
						} `json:"n_out_of"`
					} `json:"rule"`
					Version int `json:"version"`
				} `json:"value"`
			} `json:"policy"`
			Version string `json:"version"`
		} `json:"Admins"`
		Readers struct {
			ModPolicy string `json:"mod_policy"`
			Policy    struct {
				Type  int `json:"type"`
				Value struct {
					Identities []struct {
						Principal struct {
							MspIdentifier string `json:"msp_identifier"`
							Role          string `json:"role"`
						} `json:"principal"`
						PrincipalClassification string `json:"principal_classification"`
					} `json:"identities"`
					Rule struct {
						NOutOf struct {
							N     int `json:"n"`
							Rules []struct {
								SignedBy int `json:"signed_by"`
							} `json:"rules"`
						} `json:"n_out_of"`
					} `json:"rule"`
					Version int `json:"version"`
				} `json:"value"`
			} `json:"policy"`
			Version string `json:"version"`
		} `json:"Readers"`
		Writers struct {
			ModPolicy string `json:"mod_policy"`
			Policy    struct {
				Type  int `json:"type"`
				Value struct {
					Identities []struct {
						Principal struct {
							MspIdentifier string `json:"msp_identifier"`
							Role          string `json:"role"`
						} `json:"principal"`
						PrincipalClassification string `json:"principal_classification"`
					} `json:"identities"`
					Rule struct {
						NOutOf struct {
							N     int `json:"n"`
							Rules []struct {
								SignedBy int `json:"signed_by"`
							} `json:"rules"`
						} `json:"n_out_of"`
					} `json:"rule"`
					Version int `json:"version"`
				} `json:"value"`
			} `json:"policy"`
			Version string `json:"version"`
		} `json:"Writers"`
	} `json:"policies"`
	Values struct {
		MSP struct {
			ModPolicy string `json:"mod_policy"`
			Value     struct {
				Config struct {
					Admins       []string `json:"admins" validate:"required,dive,base64"`
					CryptoConfig struct {
						IdentityIdentifierHashFunction string `json:"identity_identifier_hash_function"`
						SignatureHashFamily            string `json:"signature_hash_family"`
					} `json:"crypto_config"`
					FabricNodeOus                 interface{}   `json:"fabric_node_ous"`
					IntermediateCerts             []interface{} `json:"intermediate_certs"`
					Name                          string        `json:"name"`
					OrganizationalUnitIdentifiers []interface{} `json:"organizational_unit_identifiers"`
					RevocationList                []interface{} `json:"revocation_list"`
					RootCerts                     []string      `json:"root_certs" validate:"required,dive,base64"`
					SigningIdentity               interface{}   `json:"signing_identity"`
					TLSIntermediateCerts          []interface{} `json:"tls_intermediate_certs"`
					TLSRootCerts                  []string      `json:"tls_root_certs" validate:"required,dive,base64"`
				} `json:"config"`
				Type int `json:"type"`
			} `json:"value"`
			Version string `json:"version"`
		} `json:"MSP"`
	} `json:"values"`
	Version string `json:"version"`
}
