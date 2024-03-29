package types

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	codec "github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	proto "github.com/gogo/protobuf/proto"
)

const (
	Namespace          = "ollo:"
	KeyMediaFieldValue = "value"
)

var (
	ClassKeyName             = fmt.Sprintf("%s%s", Namespace, "name")
	ClassKeySymbol           = fmt.Sprintf("%s%s", Namespace, "symbol")
	ClassKeyDescription      = fmt.Sprintf("%s%s", Namespace, "description")
	ClassKeyURIhash          = fmt.Sprintf("%s%s", Namespace, "uri_hash")
	ClassKeyMintRestricted   = fmt.Sprintf("%s%s", Namespace, "mint_restricted")
	ClassKeyUpdateRestricted = fmt.Sprintf("%s%s", Namespace, "update_restricted")
	ClassKeyCreator          = fmt.Sprintf("%s%s", Namespace, "creator")
	ClassKeySchema           = fmt.Sprintf("%s%s", Namespace, "schema")
	TokenKeyName             = fmt.Sprintf("%s%s", Namespace, "name")
	TokenKeyURIhash          = fmt.Sprintf("%s%s", Namespace, "uri_hash")
)

type (
	ClassBuilder struct {
		cdc              codec.Codec
		getModuleAddress func(string) sdk.AccAddress
	}
	MediaField struct {
		Value interface{} `json:"value"`
		Mime  string      `json:"mime,omitempty"`
	}
	TokenBuilder struct {
		cdc codec.Codec
	}
)

func NewClassBuilder(
	cdc codec.Codec,
	getModuleAddress func(string) sdk.AccAddress,
) ClassBuilder {
	return ClassBuilder{
		cdc,
		getModuleAddress,
	}
}

func (cb ClassBuilder) BuildMetadata(class nft.Class) (string, error) {
	var message proto.Message
	if err := cb.cdc.UnpackAny(class.Data, &message); err != nil {
		return "", err
	}
	metadata, ok := message.(*DenomMetadata)
	if !ok {
		return "", errors.New("unsupported classMetadata")
	}
	kvals := make(map[string]interface{})
	if len(metadata.Data) > 0 {
		if err := json.Unmarshal([]byte(metadata.Data), &kvals); err != nil &&
			IsIBCDenom(class.Id) {
			return base64.RawStdEncoding.EncodeToString([]byte(metadata.Data)), nil
		}
	}
	creator, err := sdk.AccAddressFromBech32(metadata.Creator)
	if err != nil {
		return "", err
	}

	hexCreator := hex.EncodeToString(creator)
	kvals[ClassKeyName] = MediaField{Value: class.Name}
	kvals[ClassKeySymbol] = MediaField{Value: class.Symbol}
	kvals[ClassKeyDescription] = MediaField{Value: class.Description}
	kvals[ClassKeyURIhash] = MediaField{Value: class.UriHash}
	kvals[ClassKeyMintRestricted] = MediaField{Value: metadata.MintRestricted}
	kvals[ClassKeyUpdateRestricted] = MediaField{Value: metadata.UpdateRestricted}
	kvals[ClassKeyCreator] = MediaField{Value: hexCreator}
	kvals[ClassKeySchema] = MediaField{Value: metadata.Schema}
	data, err := json.Marshal(kvals)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(data), nil
}

func (cb ClassBuilder) Build(
	classId, classURI, classData string,
) (nft.Class, error) {
	classDataBz, err := base64.RawStdEncoding.DecodeString(classData)
	if err != nil {
		return nft.Class{}, err
	}
	var (
		name             = ""
		symbol           = ""
		description      = ""
		uriHash          = ""
		mintRestricted   = true
		updateRestricted = true
		schema           = ""
		creator          = cb.getModuleAddress(ModuleName).String()
	)
	dataMap := make(map[string]interface{})
	if err := json.Unmarshal(classDataBz, &dataMap); err != nil {
		anydata, err := codectypes.NewAnyWithValue(&DenomMetadata{
			Creator:          creator,
			Schema:           schema,
			MintRestricted:   mintRestricted,
			UpdateRestricted: updateRestricted,
			Data:             string(classDataBz),
		})
		if err != nil {
			return nft.Class{}, err
		}
		return nft.Class{
			Id:          classId,
			Uri:         classURI,
			Name:        name,
			Symbol:      symbol,
			Description: description,
			UriHash:     uriHash,
			Data:        anydata,
		}, nil
	}
	if v, ok := dataMap[ClassKeyName]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vStr, ok := vMap[KeyMediaFieldValue].(string); ok {
				name = vStr
				delete(dataMap, ClassKeyName)
			}
		}
	}
	if v, ok := dataMap[ClassKeySymbol]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vStr, ok := vMap[KeyMediaFieldValue].(string); ok {
				symbol = vStr
				delete(dataMap, ClassKeySymbol)
			}
		}
	}

	if v, ok := dataMap[ClassKeyDescription]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vStr, ok := vMap[KeyMediaFieldValue].(string); ok {
				description = vStr
				delete(dataMap, ClassKeyDescription)
			}
		}
	}

	if v, ok := dataMap[ClassKeyURIhash]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vStr, ok := vMap[KeyMediaFieldValue].(string); ok {
				uriHash = vStr
				delete(dataMap, ClassKeyURIhash)
			}
		}
	}

	if v, ok := dataMap[ClassKeyMintRestricted]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vBool, ok := vMap[KeyMediaFieldValue].(bool); ok {
				mintRestricted = vBool
				delete(dataMap, ClassKeyMintRestricted)
			}
		}
	}

	if v, ok := dataMap[ClassKeyUpdateRestricted]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vBool, ok := vMap[KeyMediaFieldValue].(bool); ok {
				updateRestricted = vBool
				delete(dataMap, ClassKeyUpdateRestricted)
			}
		}
	}

	if v, ok := dataMap[ClassKeyCreator]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vStr, ok := vMap[KeyMediaFieldValue].(string); ok {
				creatorAcc, err := sdk.AccAddressFromHexUnsafe(vStr)
				if err != nil {
					return nft.Class{}, err
				}
				creator = creatorAcc.String()
				delete(dataMap, ClassKeyCreator)
			}
		}
	}

	if v, ok := dataMap[ClassKeySchema]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vStr, ok := vMap[KeyMediaFieldValue].(string); ok {
				schema = vStr
				delete(dataMap, ClassKeySchema)
			}
		}
	}
	data, err := json.Marshal(dataMap)
	if err != nil {
		return nft.Class{}, err
	}
	anydata, err := codectypes.NewAnyWithValue(&DenomMetadata{
		Creator:          creator,
		Schema:           schema,
		MintRestricted:   mintRestricted,
		UpdateRestricted: updateRestricted,
		Data:             string(data),
	})
	if err != nil {
		return nft.Class{}, err
	}
	return nft.Class{
		Id:          classId,
		Uri:         classURI,
		Name:        name,
		Symbol:      symbol,
		Description: description,
		UriHash:     uriHash,
		Data:        anydata,
	}, nil
}

func NewTokenBuilder(cdc codec.Codec) TokenBuilder {
	return TokenBuilder{cdc: cdc}
}

func (tb TokenBuilder) BuildMetadata(token nft.NFT) (string, error) {
	var msg proto.Message
	if err := tb.cdc.UnpackAny(token.Data, &msg); err != nil {
		return "", err
	}
	nftMetadata, ok := msg.(*NFTMetadata)
	if !ok {
		return "", errors.New("unsupported classMetadata")
	}
	kvals := make(map[string]interface{})
	if len(nftMetadata.Data) > 0 {
		if err := json.Unmarshal([]byte(nftMetadata.Data), &kvals); err != nil && IsIBCDenom(token.ClassId) {
			return base64.RawStdEncoding.EncodeToString([]byte(nftMetadata.Data)), nil
		}
	}
	kvals[TokenKeyName] = MediaField{Value: nftMetadata.Name}
	kvals[TokenKeyURIhash] = MediaField{Value: token.UriHash}
	data, err := json.Marshal(kvals)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(data), nil

}

func (tb TokenBuilder) Build(
	classId, tokenId, tokenURI, tokenData string,
) (nft.NFT, error) {
	tokenDataBz, err := base64.RawStdEncoding.DecodeString(tokenData)
	if err != nil {
		return nft.NFT{}, err
	}

	dataMap := make(map[string]interface{})
	if err := json.Unmarshal(tokenDataBz, &dataMap); err != nil {
		metadata, err := codectypes.NewAnyWithValue(&NFTMetadata{
			Data: string(tokenDataBz),
		})
		if err != nil {
			return nft.NFT{}, err
		}

		return nft.NFT{
			ClassId: classId,
			Id:      tokenId,
			Uri:     tokenURI,
			Data:    metadata,
		}, nil
	}

	var (
		name    string
		uriHash string
	)
	if v, ok := dataMap[TokenKeyName]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vStr, ok := vMap[KeyMediaFieldValue].(string); ok {
				name = vStr
				delete(dataMap, TokenKeyName)
			}
		}
	}
	if v, ok := dataMap[TokenKeyURIhash]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vStr, ok := vMap[KeyMediaFieldValue].(string); ok {
				uriHash = vStr
				delete(dataMap, TokenKeyURIhash)
			}
		}
	}

	data, err := json.Marshal(dataMap)
	if err != nil {
		return nft.NFT{}, err
	}

	metadata, err := codectypes.NewAnyWithValue(&NFTMetadata{
		Name: name,
		Data: string(data),
	})
	if err != nil {
		return nft.NFT{}, err
	}

	return nft.NFT{
		ClassId: classId,
		Id:      tokenId,
		Uri:     tokenURI,
		UriHash: uriHash,
		Data:    metadata,
	}, nil
}
