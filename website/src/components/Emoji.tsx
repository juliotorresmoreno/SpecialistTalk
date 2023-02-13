import React from 'react'
import styled from 'styled-components'
import emojis from '../data/emojis'
import useFormValue from '../hooks/useFormValue'
import useToggle from '../hooks/useToggle'
import Anchor from './Anchor'
import Select from './Select'

const Floating = styled.div`
  position: absolute;
  background-color: var(--bs-white);
  border: 1px solid var(--bs-gray-200);
  transform: translateY(-100%);
  padding: var(--spacing-v1);
  user-select: none;
`

const Emojis = styled.div`
  margin-top: var(--spacing-v1);
  font-size: 1.5rem;

  height: 500px;
  width: 400px;

  //gap: var(--spacing-v1);
`

const Item = styled.span``

const Emoji: React.FC = () => {
  const [select, onSelect] = useFormValue(0)
  const [isOpen, toggle] = useToggle(true)

  const items = (emojis.at(select)?.items || []).slice(0, 300)

  return (
    <>
      {isOpen ? (
        <Floating>
          <Select value={select} onChange={onSelect}>
            {emojis.map((card, key) => (
              <option key={'emoji-name-' + key} value={key}>
                {card.name}
              </option>
            ))}
          </Select>
          <Emojis>
            {items.map((item, key) => (
              <Item key={'emoji-' + key}>{item}</Item>
            ))}
          </Emojis>
        </Floating>
      ) : null}
      <Anchor onClick={toggle}>
        <span className="material-symbols-outlined">add_reaction</span>
      </Anchor>
    </>
  )
}

export default Emoji
