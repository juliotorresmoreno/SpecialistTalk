import React from 'react'
import styled from 'styled-components'
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
`

const Input = styled(_Input)`
  flex: 1;
  height: calc(var(--spacing-v1) * 3.5);
`

export const Messages = () => {
  return (
    <Container>
      <Content>sss</Content>
      <InputContainer>
        <Input />
        <Button>
          <span className="material-symbols-outlined">send</span>
        </Button>
      </InputContainer>
    </Container>
  )
}
