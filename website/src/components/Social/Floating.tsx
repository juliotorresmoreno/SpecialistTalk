import React from 'react'
import { useNavigate } from 'react-router'
import styled from 'styled-components'
import config from '../../config'
import withData from '../../hoc/withData'
import useSearch from '../../hooks/useSearch'
import { IChat } from '../../models/chat'
import { useAppSelector } from '../../store/hooks'
import { HTTPError } from '../../types/http'
import Input from '../Input'

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

const InputSearch = Input

type _FloatingProps = {
  payload: IChat[]
}

const _Floating: React.FC<_FloatingProps> = ({ payload }) => {
  const url = '/users'
  const [search, setSearch, result] = useSearch(url)
  const navigate = useNavigate()
  const session = useAppSelector((state) => state.auth.session)

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
          handler: () => {
            const url = config.baseUrl + '/chats/' + el.id
            fetch(url, {
              method: 'PUT',
              headers: {
                'X-API-Key': session?.token ?? '',
              },
            })
              .then(async (response) => {
                const content = await response.json()
                if (response.ok) {
                  const chat: IChat = content
                  navigate('/chats/' + chat.code)
                  return
                }
                throw content
              })
              .catch((err: Error | HTTPError) => {
                alert(err.message)
              })
          },
        }))
      : payload.map((el) => ({
          id: el.code,
          name: (
            <>
              <span className="material-symbols-outlined">person</span>
              &nbsp;
              {el.name}
            </>
          ),
          handler: () => {
            navigate('/chats/' + el.code)
          },
        }))

  return (
    <>
      <InputSearch type="text" value={search} onChange={setSearch} />
      <ContactsContainer>
        {users.map((contact) => (
          <Contact key={contact.id} onClick={contact.handler}>
            {contact.name}
          </Contact>
        ))}
      </ContactsContainer>
    </>
  )
}

const url = config.baseUrl + '/chats'
const Floating = withData(_Floating, url)

export default Floating
