import React from 'react'
import { Button } from 'react-bootstrap'
import styled from 'styled-components'

const Container = styled.div`
  display: flex;
  flex: auto;
  justify-content: end;
  padding: var(--spacing-v1);
  gap: var(--spacing-v1);
`
const notAction: React.MouseEventHandler<HTMLAnchorElement> = (evt) =>
  evt.preventDefault()

const Anchor = styled.a.attrs({ onClick: notAction })`
  color: var(--bs-gray-700);
`

const Actions: React.FC = () => {
  return (
    <Container>
      <Anchor href="">
        <span className="material-symbols-outlined">call</span>
      </Anchor>
      <Anchor href="">
        <span className="material-symbols-outlined">videocam</span>
      </Anchor>
      <Anchor href="">
        <span className="material-symbols-outlined">mic</span>
      </Anchor>
      <Anchor href="">
        <span className="material-symbols-outlined">attach_file_add</span>
      </Anchor>
    </Container>
  )
}

export default Actions
