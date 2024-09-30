/// <reference types="vitest/config" />
import {defineConfig} from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  base: '/app',
  test: {
    // 👋 add the line below to add jsdom to vite
    environment: 'jsdom',
    reporters: [
      'default', // Vitest's default reporter so that terminal output is still visible
      ['vitest-sonar-reporter', { outputFile: './reports/test-report.xml' }],
    ],
    coverage: {
      reporter: ['lcov', 'html'],
      reportsDirectory: './reports'
    }
  }
})
