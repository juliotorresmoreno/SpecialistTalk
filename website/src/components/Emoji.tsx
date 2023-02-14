import React, { useState } from 'react'
import { OverlayTrigger, Tooltip } from 'react-bootstrap'
import styled from 'styled-components'
import emojis from '../data/emojis'
import useToggle from '../hooks/useToggle'
import Anchor from './Anchor'
import Input from './Input'
import Paginator from './Paginator'
import Select from './Select'

const itemsPerRow = 15
const itemsPerColumn = 12

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

  display: grid;
  grid-template-columns: repeat(${itemsPerColumn}, 1fr);
  grid-template-rows: repeat(${itemsPerRow}, 1fr);
  overflow-y: visible;
`

const Item = styled.span`
  display: flex;
  flex-direction: row;
  cursor: pointer;
`
const Share = styled.div`
  display: flex;
  gap: var(--spacing-v1);
`
const _allitems = itemsPerColumn * itemsPerRow

type EmojiProps = {
  onSelect: (x: string) => void
}

const Emoji: React.FC<EmojiProps> = (props) => {
  const normalize = (
    select: number,
    page: number,
    items?: [string, string][]
  ): [number, [string, string][]] => {
    const _items = items || emojis.at(select)?.items || []

    const start = _allitems * (page - 1)
    const end = _allitems * page
    const selectItems = _items.slice(start, end)
    const pages = Math.ceil(_items.length / _allitems)

    return [pages, selectItems]
  }

  const [_pages, _selectItems] = normalize(0, 1)

  const [select, setSelect] = useState(0)
  const [isOpen, toggle] = useToggle(false)
  const [page, setPage] = useState(1)
  const [pages, setPages] = useState(_pages)
  const [selectItems, setSelectItems] = useState(_selectItems)
  const [search, setSearch] = useState('')

  const onSelect: React.ChangeEventHandler<HTMLSelectElement> = (evt) => {
    const select = parseInt(evt.target.value)
    const [pages, selectItems] = normalize(select, 1)
    setSelect(select)
    setPage(1)
    setPages(pages)
    setSelectItems(selectItems)
  }

  const _setPage = (page: number) => {
    const [_, selectItems] = normalize(select, page)
    setPage(page)
    setSelectItems(selectItems)
  }

  const _onSearch: React.ChangeEventHandler<HTMLInputElement> = (evt) => {
    evt.preventDefault()
    const search = evt.target.value

    const _search = search.toLowerCase()
    const items = (emojis.at(select)?.items || []).filter((el) => {
      return el[1].toLowerCase().includes(_search)
    })

    const [_, selectItems] = normalize(select, 1, items)
    setPage(1)
    setSearch(search)
    setSelectItems(selectItems)
  }

  const onItemSelect = (emoji: string) => {
    toggle()
    props.onSelect(emoji)
  }

  return (
    <>
      {isOpen ? (
        <Floating>
          <Share>
            <Select value={select} onChange={onSelect}>
              {emojis.map((card, key) => (
                <option key={'emoji-name-' + key} value={key}>
                  {card.name}
                </option>
              ))}
            </Select>
            <Input value={search} onChange={_onSearch} />
          </Share>
          <Emojis>
            {selectItems.map(([emoji, desc], key) => (
              <OverlayTrigger
                key={'emoji-' + key}
                placement="bottom"
                overlay={
                  <Tooltip id="tooltip">
                    <strong>{desc}</strong>
                  </Tooltip>
                }
              >
                <Item onClick={() => onItemSelect(emoji)}>{emoji}</Item>
              </OverlayTrigger>
            ))}
          </Emojis>
          <Paginator pages={pages} page={page} onChange={_setPage} />
        </Floating>
      ) : null}
      <Anchor onClick={toggle}>
        <span className="material-symbols-outlined">add_reaction</span>
      </Anchor>
    </>
  )
}

export default Emoji
