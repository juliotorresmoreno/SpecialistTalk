export default {
  title: process.env.APP_NAME,
  baseUrl: process.env.BASE_URL || document.location.href + 'api/v1',
  wsUrl: process.env.WS_URL || document.location.href + 'realtime',
}
