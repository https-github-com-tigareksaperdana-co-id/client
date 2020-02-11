import * as React from 'react'
import * as Types from '../../../constants/types/fs'
import * as Kb from '../../../common-adapters'
import * as Kbfs from '../../common'
import * as Styles from '../../../styles'
import * as FsGen from '../../../actions/fs-gen'
import * as Chat2Gen from '../../../actions/chat2-gen'
import * as ChatTypes from '../../../constants/types/chat2'
import * as Container from '../../../util/container'
import * as RouteTreeGen from '../../../actions/route-tree-gen'
import ConversationList from './conversation-list/conversation-list'
import ChooseConversation from './conversation-list/choose-conversation'

type Props = {
  onCancel: () => void
  onSetTitle: (title: string) => void
  send?: () => void
  path: Types.Path
  sendAttachmentToChatState: Types.SendAttachmentToChatState
  title: string
}

const MobileConnectedConversationList = (props: {customComponent?: React.ReactNode | null}) => {
  const dispatch = Container.useDispatch()
  const sendAttachmentToChat = Container.useSelector(state => state.fs.sendAttachmentToChat)
  const username = Container.useSelector(state => state.config.username)
  const onSelect = (conversationIDKey: ChatTypes.ConversationIDKey, convName: string) => {
    dispatch(
      Chat2Gen.createAttachmentsUpload({
        conversationIDKey,
        paths: [{outboxID: null, path: Types.pathToString(sendAttachmentToChat.path)}],
        titles: [sendAttachmentToChat.title],
        tlfName: `${username},${convName.split('#')[0]}`,
      })
    )
    dispatch(RouteTreeGen.createClearModals())
    dispatch(Chat2Gen.createSelectConversation({conversationIDKey, reason: 'files'}))
    dispatch(Chat2Gen.createNavigateToThread())
    dispatch(FsGen.createSentAttachmentToChat())
  }
  const additionalProps = {
    onSelect,
  }
  return <ConversationList {...props} {...additionalProps} />
}

const MobileWithHeader = Kb.HeaderHoc(MobileConnectedConversationList)

const MobileHeader = (props: Props) => (
  <Kb.Box2 direction="horizontal" centerChildren={true} fullWidth={true} style={mobileStyles.headerContainer}>
    <Kb.Text type="BodyBigLink" style={mobileStyles.button} onClick={props.onCancel}>
      Cancel
    </Kb.Text>
    <Kb.Box2 direction="horizontal" style={mobileStyles.headerContent} fullWidth={true} centerChildren={true}>
      <Kb.Text type="BodySemibold" style={mobileStyles.filename}>
        {Types.getPathName(props.path)}
      </Kb.Text>
    </Kb.Box2>
  </Kb.Box2>
)

const DesktopConversationDropdown = (props: {dropdownButtonStyle?: Styles.StylesCrossPlatform | null}) => {
  const dispatch = Container.useDispatch()
  const sendAttachmentToChat = Container.useSelector(state => state.fs.sendAttachmentToChat)
  const onSelect = (convID: ChatTypes.ConversationIDKey, convName: string) => {
    dispatch(FsGen.createSetSendAttachmentToChatConvID({convID, convName}))
  }
  const additionalProps = {
    onSelect,
    selected: sendAttachmentToChat.convID,
  }
  return <ChooseConversation {...props} {...additionalProps} />
}

const DesktopSendAttachmentToChat = (props: Props) => (
  <>
    <Kb.Box2 direction="vertical" style={desktopStyles.container} centerChildren={true}>
      <Kb.Box2 direction="horizontal" centerChildren={true} style={desktopStyles.header} fullWidth={true}>
        <Kb.Text type="Header">Attach in conversation</Kb.Text>
      </Kb.Box2>
      <Kb.Box2 direction="vertical" style={desktopStyles.belly} fullWidth={true}>
        <Kb.Box2
          direction="vertical"
          centerChildren={true}
          fullWidth={true}
          style={desktopStyles.pathItem}
          gap="tiny"
        >
          <Kbfs.ItemIcon size={48} path={props.path} badgeOverride="iconfont-attachment" />
          <Kb.Text type="BodySmall">{Types.getPathName(props.path)}</Kb.Text>
        </Kb.Box2>
        <DesktopConversationDropdown dropdownButtonStyle={desktopStyles.dropdown} />
        <Kb.LabeledInput
          placeholder="Title"
          value={props.title}
          style={desktopStyles.input}
          onChangeText={props.onSetTitle}
        />
      </Kb.Box2>
      <Kb.ButtonBar fullWidth={true} style={desktopStyles.buttonBar}>
        <Kb.Button type="Dim" label="Cancel" onClick={props.onCancel} />
        <Kb.Button
          label="Send in conversation"
          onClick={props.send}
          disabled={props.sendAttachmentToChatState !== Types.SendAttachmentToChatState.ReadyToSend}
        />
      </Kb.ButtonBar>
    </Kb.Box2>
  </>
)

const SendAttachmentToChat = Styles.isMobile
  ? (props: Props) => <MobileWithHeader customComponent={<MobileHeader {...props} />} />
  : Kb.HeaderOrPopup(DesktopSendAttachmentToChat)

export default SendAttachmentToChat

const mobileStyles = Styles.styleSheetCreate(
  () =>
    ({
      button: {
        paddingBottom: Styles.globalMargins.tiny,
        paddingLeft: Styles.globalMargins.small,
        paddingRight: Styles.globalMargins.small,
        paddingTop: Styles.globalMargins.tiny,
      },
      filename: {
        textAlign: 'center',
      },
      headerContainer: {
        minHeight: 44,
      },
      headerContent: {
        flex: 1,
        flexShrink: 1,
        padding: Styles.globalMargins.xtiny,
      },
    } as const)
)

const desktopStyles = Styles.styleSheetCreate(
  () =>
    ({
      belly: {
        ...Styles.globalStyles.flexGrow,
        alignItems: 'center',
        marginBottom: Styles.globalMargins.small,
        paddingLeft: Styles.globalMargins.large,
        paddingRight: Styles.globalMargins.large,
      },
      buttonBar: {alignItems: 'center'},
      container: Styles.platformStyles({
        isElectron: {
          maxHeight: 560,
          width: 400,
        },
      }),
      dropdown: {
        marginBottom: Styles.globalMargins.small,
        marginTop: Styles.globalMargins.mediumLarge,
        width: '100%',
      },
      header: {
        paddingTop: Styles.globalMargins.mediumLarge,
      },
      input: {
        width: '100%',
      },
      pathItem: {
        marginTop: Styles.globalMargins.mediumLarge,
      },
    } as const)
)
