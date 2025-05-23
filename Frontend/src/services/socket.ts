import { useNotifier } from '@/composables/useNotifier' // Импортируем useNotifier

let socket: WebSocket | null = null
let onMessageCallback: ((data: Record<string, any>) => void) | null = null

export function initSocket(url: string) {
  const { notify } = useNotifier()

  socket = new WebSocket(url)
  const token = "Bearer " + localStorage.getItem('token');

  socket.onopen = () => {
    console.log('[WebSocket] Connected')

    if (token) {
      const authMessage = JSON.stringify({
        type: 'authenticate',
        token: token,
      })
      socket?.send(authMessage)
      console.log('[WebSocket] Sent auth token')
    } else {
      console.warn('[WebSocket] No token found, skipping authentication')
    }
  }


  socket.onmessage = (event: MessageEvent) => {
    try {
      const raw = JSON.parse(event.data)
      console.log('[WebSocket] Parsed message:', raw)

      const messageText = formatSocketMessage(raw)
      notify(messageText, 'info')
    } catch (e) {
      console.error('[WebSocket] Failed to parse message:', event.data, e)
    }
  }

  socket.onclose = () => {
    console.warn('[WebSocket] Disconnected. Reconnecting...')
    setTimeout(() => initSocket(url), 3000)
  }
}

export function onMessage(callback: (data: Record<string, any>) => void) {
  onMessageCallback = callback
}

export function formatSocketMessage(data: Record<string, any>): string {
  return Object.entries(data)
    .map(([key, value]) => `${key}: ${value}`)
    .join('\n')
}
