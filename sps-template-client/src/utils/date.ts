import moment, { MomentInput } from 'moment'

export const formatTime = (time: MomentInput, format = 'YYYY-MM-DD HH:mm:ss') => {
  return moment(time).format(format)
}
