import request from '@/utils/request'
import { FileUploadAndDownloadParams } from './model/fileUploadAndDownload'
import { PageInfoParams } from './model/common'

// @Tags FileUploadAndDownload
// @Summary 上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body true "上传文件"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileUploadAndDownload/getFileList [post]
export const uploadFile = (data: FormData, token: string) => {
  return request({
    url: '/fileUploadAndDownload/upload',
    method: 'post',
    headers: {
      'x-token': token
    },
    data
  })
}

// @Tags FileUploadAndDownload
// @Summary 分页文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取文件户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileUploadAndDownload/getFileList [post]
export const getFileList = (data: PageInfoParams) => {
  return request({
    url: '/fileUploadAndDownload/getFileList',
    method: 'post',
    data
  })
}

// @Tags FileUploadAndDownload
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body dbModel.FileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /fileUploadAndDownload/deleteFile [post]
export const deleteFile = (data: FileUploadAndDownloadParams) => {
  return request({
    url: '/fileUploadAndDownload/deleteFile',
    method: 'post',
    data
  })
}
