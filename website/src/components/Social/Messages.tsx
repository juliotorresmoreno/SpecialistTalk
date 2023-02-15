import React, { useEffect, useRef } from 'react'
import { useParams } from 'react-router'
import styled from 'styled-components'
import { useAdd } from '../../services/messages'
import { useAppSelector } from '../../store/hooks'
import Button from '../Button'
import Emoji from '../Emoji'
import _Input from '../Input'
import Message from './Message'

const Container = styled.div`
  background-color: white;
  display: flex;
  flex: 1;
  flex-direction: column;
`

const Content = styled.div`
  overflow-y: scroll;
  height: calc(100vh - 163px);
`

const InputContainer = styled.div`
  background-color: white;
  display: flex;
  height: calc(var(--spacing-v1) * 3.5);
  gap: var(--spacing-v1);
`

const Input = styled(_Input)`
  flex: 1;
  height: calc(var(--spacing-v1) * 3.5);
`

const Messages = () => {
  const contentRef = useRef<HTMLDivElement>(null)
  const inputRef = useRef<HTMLInputElement>(null)
  const { isLoading, error, add } = useAdd()
  const id = useParams().id as string
  const notifications = useAppSelector((state) => state.messages.notifications)
  const onKeyUp: React.KeyboardEventHandler<HTMLInputElement> = (evt) => {
    const current = inputRef.current
    if (evt.key !== 'Enter') return
    if (!current) return

    const message = current.value
    add({
      code: id as string,
      message,
    })
    current.value = ''
  }
  const onSend: React.MouseEventHandler<HTMLButtonElement> = (evt) => {
    const current = inputRef.current
    if (!current) return

    const message = current.value
    current.value = ''
    add({
      code: id as string,
      message,
    })
  }

  useEffect(() => {
    const current = contentRef.current
    const clientHeight = current?.clientHeight ?? 0
    const scrollHeight = current?.scrollHeight ?? 0
    const top = scrollHeight - clientHeight
    current?.scrollTo({
      top: top,
    })
  })

  const messages = notifications[id]?.messages ?? []

  const onEmojiSelect = (emoji: string) => {
    const current = inputRef.current
    if (!current) return

    const message = current.value
    const selectionStart = current.selectionStart ?? 0
    const selectionEnd = current.selectionEnd ?? 0

    current.value =
      message.substring(0, selectionStart) +
      emoji +
      message.substring(selectionEnd)

    current.setSelectionRange(selectionStart + 2, selectionStart + 2)
    current.focus()
  }

  return (
    <Container>
      <Content ref={contentRef}>
        {messages.map((message, key) => (
          <Message key={'message-' + key} data={message} />
        ))}
      </Content>
      <InputContainer>
        <Emoji onSelect={onEmojiSelect} />
        <Input ref={inputRef} onKeyUp={onKeyUp} />
        <Button onClick={onSend}>
          <span className="material-symbols-outlined">send</span>
        </Button>
      </InputContainer>
    </Container>
  )
}

export default Messages
