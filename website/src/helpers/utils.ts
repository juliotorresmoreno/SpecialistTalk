import { FileBase64 } from '../models/files'

export function readFile(file: File) {
  return new Promise<FileBase64>(function (resolve, reject) {
    const reader = new FileReader()
    reader.readAsDataURL(file)
    reader.onload = function () {
      const url = URL.createObjectURL(file)

      resolve({
        name: file.name,
        body: reader.result?.toString(),
        url,
      })
    }
    reader.onerror = function (error) {
      reject(error)
    }
  })
}
