package models

type Link struct {
	UID        string   `json:"uid,omitempty"`
	LinkType   LinkType `json:"linktype,omitempty"`
	BggId      string   `json:"bggid,omitempty"`
	LinkValue  string   `json:"linkvalue,omitempty"`
	Inbound    bool     `json:"inbound,omitempty"`
	DgraphType string   `json:"dgraph.type"`
}

func (l *Link) SetLinkType(lt string) {
	l.LinkType = getLinkType(lt)
}

type LinkType string

const (
	BoardgameCategoryType       LinkType = "boardgamecategory"
	BoardgameMechanicType       LinkType = "boardgamemechanic"
	BoardgameFamilyType         LinkType = "boardgamefamily"
	BoardgameExpansionType      LinkType = "boardgameexpansion"
	BoardgameDesignerType       LinkType = "boardgamedesigner"
	BoardgameArtistType         LinkType = "boardgameartist"
	BoardgamePublisherType      LinkType = "boardgamepublisher"
	BoardgameImplementationType LinkType = "boardgameimplementation"
	BoardgameCompilationType    LinkType = "boardgamecompilation"
	LinkTypeNotRecognised       LinkType = "linktypenotrecognised"
)

func getLinkType(lt string) LinkType {
	switch lt {
	case string(BoardgameCategoryType):
		return BoardgameCategoryType
	case string(BoardgameMechanicType):
		return BoardgameMechanicType
	case string(BoardgameFamilyType):
		return BoardgameFamilyType
	case string(BoardgameExpansionType):
		return BoardgameExpansionType
	case string(BoardgameDesignerType):
		return BoardgameDesignerType
	case string(BoardgameArtistType):
		return BoardgameArtistType
	case string(BoardgamePublisherType):
		return BoardgamePublisherType
	case string(BoardgameImplementationType):
		return BoardgameImplementationType
	case string(BoardgameCompilationType):
		return BoardgameCompilationType
	default:
		return LinkTypeNotRecognised
	}
}
