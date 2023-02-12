import React from 'react'
import { useParams } from 'react-router'
import styled from 'styled-components'
import useFormValue from '../../hooks/useFormValue'
import { useAdd } from '../../services/messages'
import Button from '../Button'
import _Input from '../Input'

const Container = styled.div`
  background-color: white;
  display: flex;
  flex: 1;
  flex-direction: column;
`

const Content = styled.div`
  background-color: blue;
  display: flex;
  flex: 1;
`

const InputContainer = styled.div`
  background-color: white;
  display: flex;
  height: calc(var(--spacing-v1) * 3.5);
`

const Input = styled(_Input)`
  flex: 1;
  height: calc(var(--spacing-v1) * 3.5);
`

export const Messages = () => {
  const { isLoading, error, add } = useAdd()
  const [message, handlerMessage, setMessage] = useFormValue('')
  const { id } = useParams()
  const onKeyUp: React.KeyboardEventHandler<HTMLInputElement> = (evt) => {
    if (evt.key !== 'Enter') return

    add({
      code: id as string,
      message,
    }).then(() => setMessage(''))
  }

  return (
    <Container>
      <Content>Messages</Content>
      <InputContainer>
        <Input onChange={handlerMessage} onKeyUp={onKeyUp} value={message} />
        <Button>
          <span className="material-symbols-outlined">send</span>
        </Button>
      </InputContainer>
    </Container>
  )
}
