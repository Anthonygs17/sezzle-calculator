import type {
  CalculateRequest,
  CalculateResponse,
  ApiErrorResponse,
} from '../types/calculator';

const API_URL = '/api/calculate';

export async function calculate(
  request: CalculateRequest,
): Promise<CalculateResponse> {
  const response = await fetch(API_URL, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(request),
  });

  const data = (await response.json()) as
    | CalculateResponse
    | ApiErrorResponse;

  if (!response.ok) {
    const errorData = data as ApiErrorResponse;
    throw new Error(errorData.error || 'Failed to calculate');
  }

  return data as CalculateResponse;
}