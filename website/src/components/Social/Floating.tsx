import React, { useState } from 'react'
import { useLocation, useNavigate } from 'react-router'
import styled from 'styled-components'
import config from '../../config'
import chatsSlice from '../../features/chats'
import withData from '../../hoc/withData'
import useSearch from '../../hooks/useSearch'
import { IChat } from '../../models/chat'
import { useAdd } from '../../services/chats'
import { store } from '../../store'
import { useAppSelector } from '../../store/hooks'
import Ads from '../Ads'
import Input from '../Input'

const ContactsContainer = styled.div`
  display: flex;
  flex: 1;
  flex-direction: column;
  padding: var(--spacing-v1);
`

type ContactProps = { notification: boolean }
const Contact = styled.div<ContactProps>`
  height: auto;
  line-height: initial;
  cursor: pointer;

  white-space: nowrap;
  text-overflow: ellipsis;
  overflow: hidden;

  color: ${({ notification }) => (notification ? 'var(--bs-orange)' : 'black')};
`

const InputSearch = styled(Input)`
  height: calc(var(--spacing-v1) * 3.5);
`

type _FloatingProps = {
  payload: IChat[]
}

const _Floating: React.FC<_FloatingProps> = () => {
  const url = '/users'
  const [ignore, setIgnore] = useState<any>({})
  const [search, setSearch, result] = useSearch(url)
  const navigate = useNavigate()
  const { add } = useAdd()
  const contacts = useAppSelector((state) => state.chats.contacts)
  const location = useLocation()
  const path = location.pathname.split('?')[0].split('/')
  const code = path[1] === 'chats' ? path[2] : ''

  const users =
    result.length > 0
      ? result.map((el) => ({
          id: el.id.toString(),
          name: (
            <>
              <span className="material-symbols-outlined">person_add</span>
              &nbsp;
              {el.getFullName()}
            </>
          ),
          notifications: 0,
          handler: () => {
            add({ user_id: el.id })
              .then(async (chat) => {
                setIgnore({
                  ...ignore,
                  [chat.id]: true,
                })
                navigate('/chats/' + chat.code)
              })
              .catch((err: Error) => {})
          },
        }))
      : contacts.map((el) => ({
          id: el.code,
          name: (
            <>
              <span className="material-symbols-outlined">person</span>
              &nbsp;{el.name}
            </>
          ),
          notifications: el.notifications,
          handler: () => {
            navigate('/chats/' + el.code)
          },
        }))

  return (
    <>
      <InputSearch type="text" value={search} onChange={setSearch} />
      <ContactsContainer>
        {users.map(({ id, notifications, name, handler }) => {
          if (ignore[id]) return null
          return (
            <Contact
              notification={notifications > 0 && id !== code}
              key={'contact-' + id}
              onClick={handler}
            >
              {name}
            </Contact>
          )
        })}
      </ContactsContainer>
    </>
  )
}

type FloatingProps = {}

const url = config.baseUrl + '/chats'
const Floating = withData<FloatingProps, IChat[]>({
  WrappedComponent: _Floating,
  withAuth: true,
  url,
  FallbackComponent: Ads,
  callback(data) {
    store.dispatch(chatsSlice.actions.updateContacts(data))
  },
})

export default Floating
