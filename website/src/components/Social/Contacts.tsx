import React, { useContext } from 'react'
import styled from 'styled-components'
import useToggle from '../../hooks/useToggle'
import { User } from '../../models/user'
import Context from '../../contexts/SocialContext'
import { useLocation, useNavigate, useNavigation } from 'react-router'

const Container = styled.div`
  display: flex;
  flex: 1;
`

const ToggleButton = styled.div`
  cursor: pointer;
  flex: 1;
  text-align: center;
`

const Floating = styled.div`
  display: flex;
  flex: 1;
  flex-direction: column;
  position: absolute;
  height: calc(100vh - 106px);
  width: 100%;
  transform: translateY(-100%);
  background-color: red;
`

const InputSearch = styled.input`
  border-radius: 0;
  box-shadow: none;
  border: none;
  border-width: 0;
  width: 100%;
  -moz-box-shadow: none;
  -goog-ms-box-shadow: none;
  -webkit-box-shadow: none;
  padding: 0 var(--spacing-v1);
  &:hover,
  &:focus,
  &:not(:focus) {
    border-radius: none;
    box-shadow: none;
    border: none;
    border-width: none;
    outline: none;

    -moz-box-shadow: none;
    -goog-ms-box-shadow: none;
    -webkit-box-shadow: none;
  }
`

const ContactsContainer = styled.div`
  display: flex;
  flex: 1;
  flex-direction: column;
  padding: var(--spacing-v1);
`

const Contact = styled.div`
  height: auto;
  line-height: initial;
  cursor: pointer;

  white-space: nowrap;
  text-overflow: ellipsis;
  overflow: hidden;
`

const contacts = [
  new User({
    id: 1,
    name: 'Jhon',
    lastname: 'Doe',
    email: 'email@email.com',
    username: 'username',
  }),
  new User({
    id: 2,
    name: 'Jhon',
    lastname: 'Doe',
    email: 'email@email.com',
    username: 'username',
  }),
]

export const Contacts = () => {
  const [isOpen, toggle] = useToggle(true)
  const { setActiveChat, activeChat } = useContext(Context)

  const onToggle = () => {
    if (isOpen) setActiveChat(null)
    toggle()
  }
  const navigate = useNavigate()

  return (
    <Container>
      {isOpen ? (
        <Floating>
          <InputSearch type="text" />
          <ContactsContainer>
            {contacts.map((contact) => (
              <Contact
                key={contact.id}
                onClick={() => navigate('/chats/' + contact.id)}
              >
                {contact.getFullName()}
              </Contact>
            ))}
          </ContactsContainer>
        </Floating>
      ) : null}
      <ToggleButton onClick={onToggle}>Contactos</ToggleButton>
    </Container>
  )
}
