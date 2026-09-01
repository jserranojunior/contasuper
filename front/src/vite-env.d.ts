/// <reference types="vite/client" />

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>
  export default component
}

declare module '*.css'

/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_API_URL: string;
  readonly VITE_EMAILJS_PUBLIC_KEY: string;
  // Declare aqui outras variáveis VITE_ que você utiliza
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}