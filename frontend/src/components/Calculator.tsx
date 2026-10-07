import { useState } from 'react';
import { calculate } from '../services/calculatorApi';
import type { Operation } from '../types/calculator';

const formatResult = (value: number): string => {
  return Number.parseFloat(value.toPrecision(10)).toString();
};

export default function Calculator() {
  const [display, setDisplay] = useState('0');
  const [operand1, setOperand1] = useState<number | null>(null);
  const [operation, setOperation] = useState<Operation | null>(null);
  const [waitingForOperand, setWaitingForOperand] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');

  const handleDigit = (digit: string) => {
    setError('');

    if (waitingForOperand) {
      setDisplay(digit);
      setWaitingForOperand(false);
      return;
    }

    setDisplay((current) =>
      current === '0' ? digit : current + digit,
    );
  };

  const handleDecimal = () => {
    setError('');

    if (waitingForOperand) {
      setDisplay('0.');
      setWaitingForOperand(false);
      return;
    }

    if (!display.includes('.')) {
      setDisplay((current) => current + '.');
    }
  };

  const handleOperation = (selectedOperation: Operation) => {
    setError('');

    const currentValue = Number(display);

    if (Number.isNaN(currentValue)) {
      setError('Invalid number.');
      return;
    }

    setOperand1(currentValue);
    setOperation(selectedOperation);
    setWaitingForOperand(true);
  };

  const handleEquals = async () => {
    if (operand1 === null || operation === null) {
      return;
    }

    const operand2 = Number(display);

    if (Number.isNaN(operand2)) {
      setError('Invalid number.');
      return;
    }

    setIsLoading(true);
    setError('');

    try {
      const response = await calculate({
        operation,
        operand1,
        operand2,
      });

      setDisplay(formatResult(response.result));
      setOperand1(null);
      setOperation(null);
      setWaitingForOperand(true);
    } catch (err) {
      setError(
        err instanceof Error
          ? err.message
          : 'An unexpected error occurred.',
      );
    } finally {
      setIsLoading(false);
    }
  };

  const handleClear = () => {
    setDisplay('0');
    setOperand1(null);
    setOperation(null);
    setWaitingForOperand(false);
    setError('');
  };

  return (
    <div className="calculator">
      <div className="calculator-display" role="status" aria-live="polite">
        {display}
      </div>

      {error && (
        <div className="calculator-error" role="alert">
          {error}
        </div>
      )}

      <div className="calculator-grid">
        <button
          className="button button-clear span-three"
          onClick={handleClear}
        >
          C
        </button>

        <button
          className="button button-operation"
          onClick={() => handleOperation('divide')}
        >
          ÷
        </button>

        <button className="button" onClick={() => handleDigit('7')}>
          7
        </button>
        <button className="button" onClick={() => handleDigit('8')}>
          8
        </button>
        <button className="button" onClick={() => handleDigit('9')}>
          9
        </button>
        <button
          className="button button-operation"
          onClick={() => handleOperation('multiply')}
        >
          ×
        </button>

        <button className="button" onClick={() => handleDigit('4')}>
          4
        </button>
        <button className="button" onClick={() => handleDigit('5')}>
          5
        </button>
        <button className="button" onClick={() => handleDigit('6')}>
          6
        </button>
        <button
          className="button button-operation"
          onClick={() => handleOperation('subtract')}
        >
          −
        </button>

        <button className="button" onClick={() => handleDigit('1')}>
          1
        </button>
        <button className="button" onClick={() => handleDigit('2')}>
          2
        </button>
        <button className="button" onClick={() => handleDigit('3')}>
          3
        </button>
        <button
          className="button button-operation"
          onClick={() => handleOperation('add')}
        >
          +
        </button>

        <button
          className="button span-two"
          onClick={() => handleDigit('0')}
        >
          0
        </button>

        <button className="button" onClick={handleDecimal}>
          .
        </button>

        <button
          className="button button-equals"
          onClick={handleEquals}
          disabled={isLoading}
        >
          {isLoading ? '...' : '='}
        </button>
      </div>
    </div>
  );
}