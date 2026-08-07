import { useEffect, useState } from 'react';
import { api } from '../api';
import type { AccountsPageResponse } from '../types';

interface UseAccountsPageOptions {
  resultId?: string;
  page: number;
  pageSize: number;
  search?: string;
  status?: string;
  sort?: string;
  onlyFailure?: boolean;
}

interface UseAccountsPageResult {
  data: AccountsPageResponse | null;
  loading: boolean;
  error: string;
}

export function useAccountsPage(options: UseAccountsPageOptions): UseAccountsPageResult {
  const { resultId, page, pageSize, search, status, sort, onlyFailure } = options;
  const queryKey = [resultId ?? '', page, pageSize, search ?? '', status ?? '', sort ?? '', onlyFailure ? '1' : '0'].join('\u0000');
  const [state, setState] = useState<{
    key: string;
    data: AccountsPageResponse | null;
    error: string;
  }>({ key: '', data: null, error: '' });

  useEffect(() => {
    if (!resultId) return;

    let cancelled = false;

    void api
      .getAccountsPage({
        resultId,
        page,
        pageSize,
        search,
        status,
        sort,
        onlyFailure,
      })
      .then((response) => {
        if (cancelled) return;
        setState({ key: queryKey, data: response, error: '' });
      })
      .catch((pageError) => {
        if (cancelled) return;
        setState({
          key: queryKey,
          data: null,
          error: pageError instanceof Error ? pageError.message : '加载账户分页失败',
        });
      });

    return () => {
      cancelled = true;
    };
  }, [onlyFailure, page, pageSize, queryKey, resultId, search, sort, status]);

  if (!resultId) {
    return { data: null, loading: false, error: '' };
  }
  if (state.key !== queryKey) {
    return { data: null, loading: true, error: '' };
  }
  return { data: state.data, loading: false, error: state.error };
}
