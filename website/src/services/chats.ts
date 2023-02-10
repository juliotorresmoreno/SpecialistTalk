import React from 'react'
import config from '../config'
import { IChat } from '../models/chat'
import * as api from './api'

export function useGet(id: string) {
  const url = config.baseUrl + '/chats/'
  const { get, error, isLoading } = api.useGetData(url + id)

  return { isLoading, error, get }
}

export function useAdd() {
  const url = config.baseUrl + '/chats'
  const { error, isLoading, add } = api.useAdd<{ user_id: number }, IChat>(url)

  return { isLoading, error, add }
}
