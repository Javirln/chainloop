//
// Copyright 2023 The Chainloop Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package materials

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	schemaapi "github.com/chainloop-dev/chainloop/app/controlplane/api/workflowcontract/v1"
	api "github.com/chainloop-dev/chainloop/internal/attestation/crafter/api/attestation/v1"
	"github.com/chainloop-dev/chainloop/internal/casclient"
	"github.com/chainloop-dev/chainloop/internal/schemavalidators"
	"github.com/rs/zerolog"
)

// CSAFCrafter is a crafter for CSAF VEX, CSAF Informational Advisory, CSAF Security Incident Response, and CSAF Security Advisory
// material types. It implements the Crafter interface.
type CSAFCrafter struct {
	backend *casclient.CASBackend
	*crafterCommon
}

// NewCSAFVEXCrafter creates a new CSAF VEX crafter
func NewCSAFVEXCrafter(materialSchema *schemaapi.CraftingSchema_Material, backend *casclient.CASBackend, l *zerolog.Logger) (*CSAFCrafter, error) {
	if materialSchema.Type != schemaapi.CraftingSchema_Material_CSAF_VEX {
		return nil, fmt.Errorf("material type is not CSAF_VEX format")
	}

	return baseCSAFCrafter(materialSchema, backend, l)
}

// NewCSAFInformationalAdvisoryCrafter creates a new CSAF Informational Advisory crafter
func NewCSAFInformationalAdvisoryCrafter(materialSchema *schemaapi.CraftingSchema_Material, backend *casclient.CASBackend, l *zerolog.Logger) (*CSAFCrafter, error) {
	if materialSchema.Type != schemaapi.CraftingSchema_Material_CSAF_INFORMATIONAL_ADVISORY {
		return nil, fmt.Errorf("material type is not CSAF_INFORMATIONAL_ADVISORY format")
	}

	return baseCSAFCrafter(materialSchema, backend, l)
}

// NewCSAFSecurityIncidentResponseCrafter creates a new CSAF Security Incident Response crafter
func NewCSAFSecurityIncidentResponseCrafter(materialSchema *schemaapi.CraftingSchema_Material, backend *casclient.CASBackend, l *zerolog.Logger) (*CSAFCrafter, error) {
	if materialSchema.Type != schemaapi.CraftingSchema_Material_CSAF_SECURITY_INCIDENT_RESPONSE {
		return nil, fmt.Errorf("material type is not CSAF_SECURITY_INCIDENT_RESPONSE format")
	}

	return baseCSAFCrafter(materialSchema, backend, l)
}

// NewCSAFSecurityAdvisoryCrafter creates a new CSAF Security Advisory crafter
func NewCSAFSecurityAdvisoryCrafter(materialSchema *schemaapi.CraftingSchema_Material, backend *casclient.CASBackend, l *zerolog.Logger) (*CSAFCrafter, error) {
	if materialSchema.Type != schemaapi.CraftingSchema_Material_CSAF_SECURITY_ADVISORY {
		return nil, fmt.Errorf("material type is not CSAF_SECURITY_ADVISORY format")
	}

	return baseCSAFCrafter(materialSchema, backend, l)
}

// NewCSAFCrafter creates a new CSAF crafter
func baseCSAFCrafter(materialSchema *schemaapi.CraftingSchema_Material, backend *casclient.CASBackend, l *zerolog.Logger) (*CSAFCrafter, error) {
	return &CSAFCrafter{
		backend:       backend,
		crafterCommon: &crafterCommon{logger: l, input: materialSchema},
	}, nil
}

func (i *CSAFCrafter) Craft(ctx context.Context, filepath string) (*api.Attestation_Material, error) {
	i.logger.Debug().Str("path", filepath).Msg("decoding CSAF file")
	f, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("can't open the file: %w - %w", err, ErrInvalidMaterialType)
	}

	var v interface{}
	if err := json.Unmarshal(f, &v); err != nil {
		i.logger.Debug().Err(err).Msg("error decoding file")
		return nil, fmt.Errorf("invalid CSAF file: %w", ErrInvalidMaterialType)
	}

	// The validator will try in cascade the different schemas since CSAF specification
	// is a strict schema.
	err = schemavalidators.ValidateCSAF(v)
	if err != nil {
		i.logger.Debug().Err(err).Msgf("error decoding file: %#v", err)

		return nil, fmt.Errorf("invalid CSAF file: %w", ErrInvalidMaterialType)
	}

	return uploadAndCraft(ctx, i.input, i.backend, filepath, i.logger)
}
