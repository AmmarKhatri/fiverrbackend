package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"graph-gateway/graph/model"
	"graph-gateway/protos/account"

	"google.golang.org/genproto/googleapis/type/postaladdress"
)

// Attributes is the resolver for the attributes field.
func (r *profileSectionResponseResolver) Attributes(ctx context.Context, obj *account.ProfileSectionResponse) ([]*model.Attribute, error) {
	attributes := make([]*model.Attribute, 0)
	for key, value := range obj.Attributes {
		attributes = append(attributes, &model.Attribute{
			Key:   key,
			Value: value,
		})
	}
	return attributes, nil
}

// Address is the resolver for the address field.
func (r *updatePrivateBasicUserInfoRequestResolver) Address(ctx context.Context, obj *account.UpdatePrivateBasicUserInfoRequest, data *model.PostalAddressInput) error {
	obj.Address = &postaladdress.PostalAddress{
		Revision:           int32(data.Revision),
		RegionCode:         data.RegionCode,
		LanguageCode:       data.LanguageCode,
		PostalCode:         data.PostalCode,
		SortingCode:        data.SortingCode,
		AdministrativeArea: data.AdministrativeArea,
		Locality:           data.Locality,
		Sublocality:        data.Sublocality,
		AddressLines:       data.AddressLines,
		Recipients:         data.Recipients,
		Organization:       data.Organization,
	}
	return nil
}

// Attributes is the resolver for the attributes field.
func (r *updateProfileSectionRequestResolver) Attributes(ctx context.Context, obj *account.UpdateProfileSectionRequest, data []*model.AttributeInput) error {
	attributes := make(map[string]string)
	for _, value := range data {
		attributes[value.Key] = value.Value
	}
	return nil
}

// ProfileSectionResponse returns ProfileSectionResponseResolver implementation.
func (r *Resolver) ProfileSectionResponse() ProfileSectionResponseResolver {
	return &profileSectionResponseResolver{r}
}

// UpdatePrivateBasicUserInfoRequest returns UpdatePrivateBasicUserInfoRequestResolver implementation.
func (r *Resolver) UpdatePrivateBasicUserInfoRequest() UpdatePrivateBasicUserInfoRequestResolver {
	return &updatePrivateBasicUserInfoRequestResolver{r}
}

// UpdateProfileSectionRequest returns UpdateProfileSectionRequestResolver implementation.
func (r *Resolver) UpdateProfileSectionRequest() UpdateProfileSectionRequestResolver {
	return &updateProfileSectionRequestResolver{r}
}

type profileSectionResponseResolver struct{ *Resolver }
type updatePrivateBasicUserInfoRequestResolver struct{ *Resolver }
type updateProfileSectionRequestResolver struct{ *Resolver }