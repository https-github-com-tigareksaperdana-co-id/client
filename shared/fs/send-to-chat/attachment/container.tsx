import * as ChatGen from '../../../actions/chat2-gen'
import * as FsGen from '../../../actions/fs-gen'
import * as ChatTypes from '../../../constants/types/chat2'
import * as RPCTypes from '../../../constants/types/rpc-gen'
import * as Types from '../../../constants/types/fs'
import * as Container from '../../../util/container'
import * as RouteTreeGen from '../../../actions/route-tree-gen'
import SendAttachmentToChat from '.'

type OwnProps = Container.RouteProps<{}>

const mapDispatchToProps = (dispatch: Container.TypedDispatch) => ({
  _send: (
    conversationIDKey: ChatTypes.ConversationIDKey,
    source: Types.Path | string | Array<RPCTypes.IncomingShareItem>,
    title: string
  ) => {
    dispatch(
      ChatGen.createAttachmentsUpload({
        conversationIDKey,
        paths: Array.isArray(source)
          ? source.map(item => ({
              outboxID: null,
              path: item.payloadPath,
            }))
          : [{outboxID: null, path: Types.pathToString(source)}],
        titles: [title],
      })
    )
    dispatch(RouteTreeGen.createClearModals())
    dispatch(ChatGen.createSelectConversation({conversationIDKey, reason: 'files'}))
    dispatch(ChatGen.createNavigateToThread())
    dispatch(FsGen.createSentAttachmentToChat())
  },
  onCancel: () => dispatch(RouteTreeGen.createClearModals()),
  onSetTitle: (title: string) => dispatch(FsGen.createSetSendAttachmentToChatTitle({title})),
})

export default Container.namedConnect(
  (state: Container.TypedState) => ({_sendAttachmentToChat: state.fs.sendAttachmentToChat}),
  mapDispatchToProps,
  (stateProps, dispatchProps, _: OwnProps) => {
    const {onCancel, onSetTitle} = dispatchProps
    const {_sendAttachmentToChat} = stateProps

    return {
      onCancel,
      onSetTitle,
      send: () =>
        dispatchProps._send(
          stateProps._sendAttachmentToChat.convID,
          stateProps._sendAttachmentToChat.source,
          stateProps._sendAttachmentToChat.title
        ),
      sendAttachmentToChatState: stateProps._sendAttachmentToChat.state,
      source: _sendAttachmentToChat.source,
      title: stateProps._sendAttachmentToChat.title,
    }
  },
  'SendAttachmentToChat'
)(SendAttachmentToChat)
