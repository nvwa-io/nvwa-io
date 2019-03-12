
export function formatMsgTime(dateStr) {
  const timespan = Date.parse(dateStr.replace(/-/gi, '/'))
  var dateTime = new Date(timespan)

  var year = dateTime.getFullYear()
  var month = dateTime.getMonth() + 1
  var day = dateTime.getDate()
  var hour = dateTime.getHours()
  var minute = dateTime.getMinutes()
  var second = dateTime.getSeconds()
  var now = new Date()
  var now_new = Date.parse(now.toDateString())

  var milliseconds = 0
  var timeSpanStr

  milliseconds = now_new - timespan

  if (milliseconds <= 1000 * 60 * 1) {
    timeSpanStr = '刚刚'
  } else if (1000 * 60 < milliseconds && milliseconds <= 1000 * 60 * 60) {
    timeSpanStr = Math.round((milliseconds / (1000 * 60))) + '分钟前'
  } else if (1000 * 60 * 60 < milliseconds && milliseconds <= 1000 * 60 * 60 * 24) {
    timeSpanStr = Math.round(milliseconds / (1000 * 60 * 60)) + '小时前'
  } else if (1000 * 60 * 60 * 24 < milliseconds && milliseconds <= 1000 * 60 * 60 * 24 * 3) {
    timeSpanStr = Math.round(milliseconds / (1000 * 60 * 60 * 24)) + '天前'
  } else {
    timeSpanStr = year + '-' + month + '-' + day + ' ' + hour + ':' + minute + ':' + second
  }
  return timeSpanStr
}
