import * as React from 'react'
import {ConversationIDKey} from '../../constants/types/chat2'
import * as RPCChatTypes from '../../constants/types/rpc-chat-gen'

export type RowItemSmall = {
  type: 'small'
  teamname: string
  isTeam: boolean
  conversationIDKey: ConversationIDKey
  time: number
  snippet?: string
  snippetDecoration: RPCChatTypes.SnippetDecoration
}
export type RowItemBigTeamsLabel = {
  type: 'bigTeamsLabel'
  isFiltered: boolean
  isTeam?: boolean
  teamname?: never
  conversationIDKey?: never
  snippet?: string
  snippetDecoration: RPCChatTypes.SnippetDecoration
  time?: number
}
export type RowItemBigHeader = {
  type: 'bigHeader'
  isTeam?: boolean
  teamname: string
  teamID: string
  conversationIDKey?: never
  snippet?: string
  snippetDecoration: RPCChatTypes.SnippetDecoration
  time?: number
}
export type RowItemBig = {
  type: 'big'
  isTeam?: boolean
  conversationIDKey: ConversationIDKey
  teamname: string
  channelname: string
  snippet?: string
  snippetDecoration: RPCChatTypes.SnippetDecoration
  time?: number
}
export type RowItemDivider = {
  conversationIDKey?: never
  teamname?: never
  showButton: boolean
  type: 'divider'
}
export type RowItemTeamBuilder = {
  conversationIDKey?: never
  teamname?: never
  type: 'teamBuilder'
}

export type RowItem =
  | RowItemSmall
  | RowItemBigTeamsLabel
  | RowItemBigHeader
  | RowItemBig
  | RowItemDivider
  | RowItemTeamBuilder

export type Props = {
  allowShowFloatingButton: boolean
  inboxNumSmallRows: number
  hasBigTeams: boolean
  isLoading: boolean
  isSearching: boolean
  navKey: string
  neverLoaded: boolean
  onNewChat: () => void
  onUntrustedInboxVisible: (conversationIDKeys: Array<ConversationIDKey>) => void
  rows: Array<RowItem>
  setInboxNumSmallRows: (rows: number) => void
  smallTeamsExpanded: boolean
  toggleSmallTeamsExpanded: () => void
  unreadIndices: Array<number>
}

export default class Inbox extends React.Component<Props> {}
