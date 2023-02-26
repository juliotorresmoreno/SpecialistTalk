import React from 'react'
import config from '../config'
import { IChat } from '../models/chat'
import { FileBase64 } from '../models/files'
import * as api from './api'

export function useGet(code: string) {
  const url = config.baseUrl + '/messages/'
  const { get, error, isLoading } = api.useGetData(url + code)

  return { isLoading, error, get }
}

type Chat = Omit<IChat, 'messages'>
type Attachments = Required<Omit<FileBase64, 'url'>>

type Add = {
  code: string
  message: string
  attachments?: Attachments[]
}

export function useAdd() {
  const url = config.baseUrl + '/messages'
  const { error, isLoading, add } = api.useAdd<Add, Chat>(url)

  return { isLoading, error, add }
}
