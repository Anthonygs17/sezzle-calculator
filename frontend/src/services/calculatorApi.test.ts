import { afterEach, describe, expect, it, vi } from 'vitest';

import { calculate } from './calculatorApi';

describe('calculate', () => {
  afterEach(() => {
    vi.restoreAllMocks();
  });

  it('returns the calculation result when the request succeeds', async () => {
    vi.spyOn(globalThis, 'fetch').mockResolvedValue(
      new Response(
        JSON.stringify({
          result: 15,
        }),
        {
          status: 200,
          headers: {
            'Content-Type': 'application/json',
          },
        },
      ),
    );

    const result = await calculate({
      operation: 'add',
      operand1: 10,
      operand2: 5,
    });

    expect(fetch).toHaveBeenCalledWith('/api/calculate', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        operation: 'add',
        operand1: 10,
        operand2: 5,
      }),
    });

    expect(result).toEqual({
      result: 15,
    });
  });

  it('throws the backend error message when the request fails', async () => {
    vi.spyOn(globalThis, 'fetch').mockResolvedValue(
      new Response(
        JSON.stringify({
          error: 'division by zero is not allowed',
        }),
        {
          status: 400,
          headers: {
            'Content-Type': 'application/json',
          },
        },
      ),
    );

    await expect(
      calculate({
        operation: 'divide',
        operand1: 10,
        operand2: 0,
      }),
    ).rejects.toThrow('division by zero is not allowed');
  });

  it('throws a fallback error when the backend does not provide a message', async () => {
    vi.spyOn(globalThis, 'fetch').mockResolvedValue(
      new Response(
        JSON.stringify({
          error: '',
        }),
        {
          status: 500,
          headers: {
            'Content-Type': 'application/json',
          },
        },
      ),
    );

    await expect(
      calculate({
        operation: 'add',
        operand1: 10,
        operand2: 5,
      }),
    ).rejects.toThrow('Failed to calculate');
  });
});