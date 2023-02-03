import React from 'react'
import useToggle from '../../hooks/useToggle'
import { Label } from './styles'
import { User } from '../../models/user'
import Context from './context'
import {
  Contact,
  ContactContainer,
  Floating as FloatingBS,
  FormControl,
} from './styles'

const Contacts = () => {
  const [isOpen, toggle] = useToggle()
  const { toggleChats, setActiveChat } = React.useContext(Context)
  const contacts = [
    new User({
      id: 0,
      name: 'julio cesar j. r.',
      lastname: 'torres moreno',
      email: 'sssssss@asdas.com',
      username: 'username',
    }),
  ]

  return (
    <>
      {isOpen ? (
        <FloatingBS>
          <FormControl type="text" placeholder="Search" />
          <ContactContainer>
            {contacts.map((contact) => (
              <Contact
                key={contact.id}
                onClick={() => {
                  toggleChats(contact)
                  setActiveChat(contact)
                }}
              >
                {contact.getFullName()}
              </Contact>
            ))}
          </ContactContainer>
        </FloatingBS>
      ) : null}
      <div className="contacts">
        <Label onClick={toggle}>chats</Label>
      </div>
    </>
  )
}

export default Contacts
