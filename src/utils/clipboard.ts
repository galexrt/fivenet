import { useClipboard } from '@vueuse/core';
import { copyToClipboard } from '~/composables/nui';

export function copyToClipboardWrapper(text: string): Promise<void> {
    if (isNUIAvailable()) {
        return copyToClipboard(text);
    } else {
        const clipboard = useClipboard();
        return clipboard.copy(text);
    }
}
