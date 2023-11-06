// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package generated

import (
	"fmt"
	converter "github.com/opendatahub-io/model-registry/internal/converter"
	proto "github.com/opendatahub-io/model-registry/internal/ml_metadata/proto"
	openapi "github.com/opendatahub-io/model-registry/internal/model/openapi"
)

type OpenAPIToMLMDConverterImpl struct{}

func (c *OpenAPIToMLMDConverterImpl) ConvertModelArtifact(source *converter.OpenAPIModelWrapper[openapi.ModelArtifact]) (*proto.Artifact, error) {
	var pProtoArtifact *proto.Artifact
	if source != nil {
		var protoArtifact proto.Artifact
		var pString *string
		if (*source).Model != nil {
			pString = (*source).Model.Id
		}
		pInt64, err := converter.StringToInt64(pString)
		if err != nil {
			return nil, fmt.Errorf("error setting field Id: %w", err)
		}
		protoArtifact.Id = pInt64
		protoArtifact.Name = converter.MapModelArtifactName(source)
		pInt642 := (*source).TypeId
		protoArtifact.TypeId = &pInt642
		protoArtifact.Type = converter.MapModelArtifactType((*source).Model)
		var pString2 *string
		if (*source).Model != nil {
			pString2 = (*source).Model.Uri
		}
		var pString3 *string
		if pString2 != nil {
			xstring := *pString2
			pString3 = &xstring
		}
		protoArtifact.Uri = pString3
		var pString4 *string
		if (*source).Model != nil {
			pString4 = (*source).Model.ExternalID
		}
		var pString5 *string
		if pString4 != nil {
			xstring2 := *pString4
			pString5 = &xstring2
		}
		protoArtifact.ExternalId = pString5
		mapStringPProtoValue, err := converter.MapModelArtifactProperties((*source).Model)
		if err != nil {
			return nil, fmt.Errorf("error setting field Properties: %w", err)
		}
		protoArtifact.Properties = mapStringPProtoValue
		var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
		if (*source).Model != nil {
			pMapStringOpenapiMetadataValue = (*source).Model.CustomProperties
		}
		mapStringPProtoValue2, err := converter.MapOpenAPICustomProperties(pMapStringOpenapiMetadataValue)
		if err != nil {
			return nil, fmt.Errorf("error setting field CustomProperties: %w", err)
		}
		protoArtifact.CustomProperties = mapStringPProtoValue2
		var pOpenapiArtifactState *openapi.ArtifactState
		if (*source).Model != nil {
			pOpenapiArtifactState = (*source).Model.State
		}
		protoArtifact.State = converter.MapOpenAPIModelArtifactState(pOpenapiArtifactState)
		pProtoArtifact = &protoArtifact
	}
	return pProtoArtifact, nil
}
func (c *OpenAPIToMLMDConverterImpl) ConvertModelVersion(source *converter.OpenAPIModelWrapper[openapi.ModelVersion]) (*proto.Context, error) {
	var pProtoContext *proto.Context
	if source != nil {
		var protoContext proto.Context
		var pString *string
		if (*source).Model != nil {
			pString = (*source).Model.Id
		}
		pInt64, err := converter.StringToInt64(pString)
		if err != nil {
			return nil, fmt.Errorf("error setting field Id: %w", err)
		}
		protoContext.Id = pInt64
		protoContext.Name = converter.MapModelVersionName(source)
		pInt642 := (*source).TypeId
		protoContext.TypeId = &pInt642
		protoContext.Type = converter.MapModelVersionType((*source).Model)
		var pString2 *string
		if (*source).Model != nil {
			pString2 = (*source).Model.ExternalID
		}
		var pString3 *string
		if pString2 != nil {
			xstring := *pString2
			pString3 = &xstring
		}
		protoContext.ExternalId = pString3
		mapStringPProtoValue, err := converter.MapModelVersionProperties(source)
		if err != nil {
			return nil, fmt.Errorf("error setting field Properties: %w", err)
		}
		protoContext.Properties = mapStringPProtoValue
		var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
		if (*source).Model != nil {
			pMapStringOpenapiMetadataValue = (*source).Model.CustomProperties
		}
		mapStringPProtoValue2, err := converter.MapOpenAPICustomProperties(pMapStringOpenapiMetadataValue)
		if err != nil {
			return nil, fmt.Errorf("error setting field CustomProperties: %w", err)
		}
		protoContext.CustomProperties = mapStringPProtoValue2
		pProtoContext = &protoContext
	}
	return pProtoContext, nil
}
func (c *OpenAPIToMLMDConverterImpl) ConvertRegisteredModel(source *converter.OpenAPIModelWrapper[openapi.RegisteredModel]) (*proto.Context, error) {
	var pProtoContext *proto.Context
	if source != nil {
		var protoContext proto.Context
		var pString *string
		if (*source).Model != nil {
			pString = (*source).Model.Id
		}
		pInt64, err := converter.StringToInt64(pString)
		if err != nil {
			return nil, fmt.Errorf("error setting field Id: %w", err)
		}
		protoContext.Id = pInt64
		var pString2 *string
		if (*source).Model != nil {
			pString2 = (*source).Model.Name
		}
		var pString3 *string
		if pString2 != nil {
			xstring := *pString2
			pString3 = &xstring
		}
		protoContext.Name = pString3
		pInt642 := (*source).TypeId
		protoContext.TypeId = &pInt642
		protoContext.Type = converter.MapRegisteredModelType((*source).Model)
		var pString4 *string
		if (*source).Model != nil {
			pString4 = (*source).Model.ExternalID
		}
		var pString5 *string
		if pString4 != nil {
			xstring2 := *pString4
			pString5 = &xstring2
		}
		protoContext.ExternalId = pString5
		mapStringPProtoValue, err := converter.MapRegisteredModelProperties((*source).Model)
		if err != nil {
			return nil, fmt.Errorf("error setting field Properties: %w", err)
		}
		protoContext.Properties = mapStringPProtoValue
		var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
		if (*source).Model != nil {
			pMapStringOpenapiMetadataValue = (*source).Model.CustomProperties
		}
		mapStringPProtoValue2, err := converter.MapOpenAPICustomProperties(pMapStringOpenapiMetadataValue)
		if err != nil {
			return nil, fmt.Errorf("error setting field CustomProperties: %w", err)
		}
		protoContext.CustomProperties = mapStringPProtoValue2
		pProtoContext = &protoContext
	}
	return pProtoContext, nil
}
