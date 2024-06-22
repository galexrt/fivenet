import { STRATEGIES } from 'vue-i18n-routing';

const appVersion: string = process.env.COMMIT_REF || 'COMMIT_REF';

const globalMdcComponents = ['PartialsCitizensCitizenInfoPopover', 'UAlert', 'UButton', 'UIcon'] as const;

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    srcDir: 'src/',
    dir: {
        app: 'app',
        public: 'src/public',
    },
    telemetry: false,
    ssr: false,
    extends: [process.env.NUXT_UI_PRO_PATH || '@nuxt/ui-pro'],
    modules: [
        '@nuxt/content',
        '@nuxtjs/mdc',
        '@nuxt/ui',
        'nuxt-typed-router',
        '@nuxtjs/robots',
        '@nuxt/fonts',
        '@vueuse/nuxt',
        '@pinia/nuxt',
        '@pinia-plugin-persistedstate/nuxt',
        'nuxt-zod-i18n',
        '@nuxtjs/i18n',
        'nuxt-update',
        '@nuxt/image',
    ],
    future: {
        compatibilityVersion: 4,
    },
    experimental: {
        defaults: {
            useAsyncData: {
                deep: true,
            },
        },
    },
    ui: {
        icons: ['simple-icons', 'mdi', 'flagpack'],
        safelistColors: ['primary', 'gray', 'red', 'orange', 'green', 'error', 'warn', 'info', 'success'],
    },
    fonts: {
        families: [{ name: 'DM Sans', weights: [100, 200, 300, 400, 500, 600, 700, 800, 900], global: true }],
    },
    app: {
        head: {
            charset: 'utf-8',
            viewport: 'width=device-width, initial-scale=1',
            link: [{ rel: 'icon', type: 'image/png', href: '/images/logo.png' }],
            bodyAttrs: {
                class: 'h-full',
            },
        },
        pageTransition: { name: 'page', mode: 'out-in' },
    },
    colorMode: {
        preference: 'dark',
    },
    typescript: {
        strict: true,
        tsConfig: {
            compilerOptions: {
                removeComments: true,
            },
        },
    },
    imports: {
        presets: [
            {
                from: '@protobuf-ts/runtime-rpc',
                imports: [
                    {
                        name: 'RpcError',
                        type: true,
                    },
                ],
            },
        ],
    },
    sourcemap: {
        client: true,
    },
    devtools: {
        enabled: true,
    },
    vite: {
        define: {
            APP_VERSION: `"${appVersion}"`,
        },
        build: {
            terserOptions: {
                compress: {
                    drop_console: true,
                    drop_debugger: true,
                },
            },
        },
        server: {
            proxy: {
                '/api/icons': {
                    target: 'https://api.iconify.design',
                    rewrite: (path) => path.replace(/^\/api\/icons/, ''),
                    changeOrigin: true,
                },
                '/api/grpc': {
                    target: 'http://localhost:8080',
                    changeOrigin: true,
                    ws: true,
                    proxyTimeout: 60 * 60 * 1000,
                    timeout: 60 * 60 * 1000,
                },
                '/api': 'http://localhost:8080',
            },
        },
    },
    robots: {
        rules: {
            UserAgent: '*',
            Disallow: '/',
            Allow: ['/$', '/index.html'],
        },
    },
    i18n: {
        vueI18n: './i18n.config.ts',
        strategy: STRATEGIES.NO_PREFIX,
        detectBrowserLanguage: false,
        skipSettingLocaleOnNavigate: true,
        locales: [
            {
                name: 'English',
                dir: 'ltr',
                isCatchallLocale: true,
                code: 'en',
                iso: 'en',
                files: ['en/en.json', 'en/zod.json'],
                icon: 'i-flagpack-gb-ukm',
            },
            {
                name: 'German',
                code: 'de',
                iso: 'de',
                files: ['de/de.json', 'de/zod.json'],
                icon: 'i-flagpack-de',
            },
        ],
        lazy: true,
        langDir: './lang',
        defaultLocale: 'de',
        defaultDirection: 'ltr',
        baseUrl: '',
        trailingSlash: false,
        compilation: {
            strictMessage: false,
        },
        parallelPlugin: true,
    },
    zodI18n: {
        useModuleLocale: true,
    },
    piniaPersistedstate: {
        storage: 'localStorage',
    },
    update: {
        version: appVersion,
        checkInterval: 115,
        path: '/api/version',
    },
    hooks: {
        // Define `@nuxt/ui` components as global to use them in `.md` (feel free to add those you need)
        'components:extend': (components) => {
            const globals = components.filter((c) => globalMdcComponents.includes(c.pascalName));

            globals.forEach((c) => (c.global = true));
        },
    },
});
