import {afterEach, describe, expect, test, vi} from 'vitest'
import {render, screen} from '@testing-library/react'
import App from './App'

describe('App', () => {
    afterEach(() => {
        vi.resetAllMocks()
    })

    test('App renders', async () => {
        vi.mock('./api', () => ({
            Api: vi.fn().mockImplementation(() => {
                return {
                    api: vi.fn().mockImplementation(() => {
                        return {
                            helloGet: vi.fn().mockResolvedValue({
                                data: {
                                    message: 'Hello'
                                }
                            })
                        }
                    })
                }
            })
        }))

        render(<App/>)

        await vi.waitFor(() => {
            const link = screen.getByText(/Hello/i)
            expect(link).toBeDefined()
        })
    })
})
