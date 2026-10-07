import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { beforeEach, describe, expect, it, vi } from 'vitest';

import Calculator from './Calculator';
import { calculate } from '../services/calculatorApi';

vi.mock('../services/calculatorApi', () => ({
  calculate: vi.fn(),
}));

const mockedCalculate = vi.mocked(calculate);

describe('Calculator', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('renders zero initially', () => {
    render(<Calculator />);

    expect(screen.getByRole('status')).toHaveTextContent('0');
  });

  it('allows entering numbers', async () => {
    const user = userEvent.setup();

    render(<Calculator />);

    await user.click(screen.getByRole('button', { name: '1' }));
    await user.click(screen.getByRole('button', { name: '2' }));
    await user.click(screen.getByRole('button', { name: '3' }));

    expect(screen.getByRole('status')).toHaveTextContent('123');
  });

  it('performs addition using the backend service', async () => {
    const user = userEvent.setup();

    mockedCalculate.mockResolvedValue({
      result: 5,
    });

    render(<Calculator />);

    await user.click(screen.getByRole('button', { name: '2' }));
    await user.click(screen.getByRole('button', { name: '+' }));
    await user.click(screen.getByRole('button', { name: '3' }));
    await user.click(screen.getByRole('button', { name: '=' }));

    expect(mockedCalculate).toHaveBeenCalledWith({
      operation: 'add',
      operand1: 2,
      operand2: 3,
    });

    await waitFor(() => {
      expect(screen.getByRole('status')).toHaveTextContent('5');
    });
  });

  it('clears the calculator state', async () => {
    const user = userEvent.setup();

    render(<Calculator />);

    await user.click(screen.getByRole('button', { name: '7' }));

    expect(screen.getByRole('status')).toHaveTextContent('7');

    await user.click(screen.getByRole('button', { name: 'C' }));

    expect(screen.getByRole('status')).toHaveTextContent('0');
  });

  it('handles decimal values', async () => {
    const user = userEvent.setup();

    render(<Calculator />);

    await user.click(screen.getByRole('button', { name: '1' }));
    await user.click(screen.getByRole('button', { name: '.' }));
    await user.click(screen.getByRole('button', { name: '5' }));

    expect(screen.getByRole('status')).toHaveTextContent('1.5');
  });

  it('displays backend errors', async () => {
    const user = userEvent.setup();

    mockedCalculate.mockRejectedValue(
      new Error('division by zero is not allowed'),
    );

    render(<Calculator />);

    await user.click(screen.getByRole('button', { name: '1' }));
    await user.click(screen.getByRole('button', { name: '÷' }));
    await user.click(screen.getByRole('button', { name: '0' }));
    await user.click(screen.getByRole('button', { name: '=' }));

    expect(
      await screen.findByRole('alert'),
    ).toHaveTextContent('division by zero is not allowed');
  });
});