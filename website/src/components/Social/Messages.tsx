import React, { useEffect, useRef } from 'react'
import { useParams } from 'react-router'
import styled from 'styled-components'
import useFormValue from '../../hooks/useFormValue'
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
  const { isLoading, error, add } = useAdd()
  const [message, handlerMessage, setMessage] = useFormValue('')
  const id = useParams().id as string
  const notifications = useAppSelector((state) => state.messages.notifications)
  const onKeyUp: React.KeyboardEventHandler<HTMLInputElement> = (evt) => {
    if (evt.key !== 'Enter') return
    setMessage('')

    add({
      code: id as string,
      message,
    })
  }
  const onSend: React.MouseEventHandler<HTMLButtonElement> = (evt) => {
    setMessage('')

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

  return (
    <Container>
      <Content ref={contentRef}>
        {messages.map((message, key) => (
          <Message key={'message' + key} data={message} />
        ))}
      </Content>
      <InputContainer>
        <Emoji />
        <Input onChange={handlerMessage} onKeyUp={onKeyUp} value={message} />
        <Button onClick={onSend}>
          <span className="material-symbols-outlined">send</span>
        </Button>
      </InputContainer>
    </Container>
  )
}

export default Messages
