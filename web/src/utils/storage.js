const USER_STORAGE_KEY = 'blink_user'
const WORKSPACE_STORAGE_KEY = 'blink_workspace'

export const loadJSON = (key, fallback = null) => {
  const stored = localStorage.getItem(key)
  return stored ? JSON.parse(stored) : fallback
}

export const loadUserFromStorage = () => loadJSON(USER_STORAGE_KEY)

export const saveUserToStorage = (userData) => {
  if (userData) {
    localStorage.setItem(USER_STORAGE_KEY, JSON.stringify(userData))
  }
}

export const removeUserFromStorage = () => localStorage.removeItem(USER_STORAGE_KEY)

export const loadWorkspaceFromStorage = () => loadJSON(WORKSPACE_STORAGE_KEY)

export const saveWorkspaceToStorage = (workspace) => {
  if (workspace) {
    localStorage.setItem(WORKSPACE_STORAGE_KEY, JSON.stringify(workspace))
  }
}

export const removeWorkspaceFromStorage = () => localStorage.removeItem(WORKSPACE_STORAGE_KEY)
