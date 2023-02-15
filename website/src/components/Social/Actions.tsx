import React, { useContext, useRef, useState } from 'react'
import styled from 'styled-components'
import SocialContext from './SocialContext'
import { FileBase64 } from '../../models/files'
import Anchor from '../Anchor'
import { readFile } from '../../helpers/utils'

const InputFile = styled.input.attrs((props) => {
  return {
    ...props,
    type: 'file',
    multiple: true,
  }
})`
  display: none;
`

const Container = styled.div`
  display: flex;
  flex: auto;
  justify-content: end;
  padding: var(--spacing-v1);
  gap: var(--spacing-v1);
`

const Actions: React.FC = () => {
  const { setAttachments } = useContext(SocialContext)
  const fileRef = useRef<HTMLInputElement>(null)

  const onFileClick: React.MouseEventHandler<HTMLAnchorElement> = (evt) => {
    evt.preventDefault()
    const current = fileRef.current
    if (!current) return

    current.click()
  }

  const onFileChange: React.ChangeEventHandler<HTMLInputElement> = async (
    evt
  ) => {
    evt.preventDefault()
    let files = evt.target.files
    if (!files) return

    let promises: Promise<FileBase64>[] = []
    for (let i = 0; i < files.length; i++) {
      promises.push(readFile(files[i]))
    }

    let responses = await Promise.all(promises)
    responses = responses.filter((el) => el.body)

    setAttachments(responses)
  }

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
      <Anchor href="" onClick={onFileClick}>
        <span className="material-symbols-outlined">attach_file_add</span>
      </Anchor>

      <InputFile
        ref={fileRef}
        type="file"
        multiple={true}
        onChange={onFileChange}
        style={{ display: 'none' }}
      />
    </Container>
  )
}

export default Actions
