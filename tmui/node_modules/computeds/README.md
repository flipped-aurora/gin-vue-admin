The minimal Signals implementation based on https://github.com/vuejs/core/pull/5912.

Removed readonly, writeable computed, deep ref etc.

This library is used in place of `@vue/reactivity` in the `vuejs/language-tools` repo to solve excessive computed recalculations, and prevent `vuejs/language-tools` development from being hindered by the progress of Vue core repo.
