package models

type Link struct {
	UID      string   `json:"uid"`
	LinkType LinkType `json:"linktype"`
	BggID    string   `json:"bggid"`
	Value    string   `json:"value"`
	Inbound  bool     `json:"inbound,omitempty"`
}

func NewLink() Link {
	return Link{}
}

func (l *Link) SetLinkType(lt string) {
	l.LinkType = getLinkType(lt)
}

func (l *Link) SetId(id string) {
	l.BggID = id
}

func (l *Link) SetValue(val string) {
	l.Value = val
}

func (l *Link) SetInbound(inb bool) {
	l.Inbound = inb
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
	default:
		return LinkTypeNotRecognised
	}
}
