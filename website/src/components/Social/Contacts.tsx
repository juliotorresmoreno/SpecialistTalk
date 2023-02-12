import React, { useContext } from 'react'
import styled from 'styled-components'
import useToggle from '../../hooks/useToggle'
import Context from '../../contexts/SocialContext'
import Floating from './Floating'

const Container = styled.div`
  display: flex;
  user-select: none;
`

const ToggleButton = styled.div`
  height: calc(var(--spacing-v1) * 3.5);
  cursor: pointer;
  text-align: center;
  flex: 1;
  color: white;
  background-color: var(--bs-primary);
  display: flex;
  place-items: center;
  place-content: center;
`

const FloatingContainer = styled.div`
  display: flex;
  flex-direction: column;
  position: absolute;
  height: calc(100vh - 106px);
  width: 200px;
  transform: translateY(-100%);
  background-color: var(--bs-white);
`

type ContactsProps = {}

const Contacts: React.FC<ContactsProps> = () => {
  const [isOpen, toggle] = useToggle(false)
  const { setActiveChat } = useContext(Context)

  const onToggle = () => {
    if (isOpen) setActiveChat(null)
    toggle()
  }

  return (
    <Container>
      {isOpen ? (
        <FloatingContainer>
          <Floating />
        </FloatingContainer>
      ) : null}
      <ToggleButton onClick={onToggle}>Contactos</ToggleButton>
    </Container>
  )
}

export default Contacts
