import axios from 'axios';

const API_URL = 'http://localhost:8008/api/dashboard';

let authToken = null;

export const setWorkspaceToken = (token) => {
  authToken = token;
};

function headers() {
  return {
    Authorization: `Bearer ${authToken}`,
    'Content-Type': 'application/json'
  };
}

// ──── Dashboard ────
export const dashboardApi = {
  async getMetrics() {
    const res = await axios.get(`${API_URL}`, { headers: headers() });
    return res.data.data;
  },

  async search(query) {
    const res = await axios.get(`${API_URL}/search?q=${encodeURIComponent(query)}`, { headers: headers() });
    return res.data.data;
  },

  async getNotifications() {
    const res = await axios.get(`${API_URL}/notifications`, { headers: headers() });
    return res.data;
  },

  async markNotificationRead(id) {
    const res = await axios.put(`${API_URL}/notifications/${id}/read`, {}, { headers: headers() });
    return res.data;
  },

  async markAllNotificationsRead() {
    const res = await axios.put(`${API_URL}/notifications/read-all`, {}, { headers: headers() });
    return res.data;
  }
};

// ──── Planner ────
export const plannerApi = {
  async getSessions(weekDate) {
    const params = weekDate ? `?week=${weekDate}` : '';
    const res = await axios.get(`${API_URL}/planner${params}`, { headers: headers() });
    return res.data.data;
  },

  async getSession(id) {
    const res = await axios.get(`${API_URL}/planner/${id}`, { headers: headers() });
    return res.data.data;
  },

  async createSession(data) {
    const res = await axios.post(`${API_URL}/planner`, data, { headers: headers() });
    return res.data.data;
  },

  async updateSession(id, data) {
    const res = await axios.put(`${API_URL}/planner/${id}`, data, { headers: headers() });
    return res.data.data;
  },

  async deleteSession(id) {
    await axios.delete(`${API_URL}/planner/${id}`, { headers: headers() });
  },

  async completeSession(id) {
    const res = await axios.patch(`${API_URL}/planner/${id}/complete`, {}, { headers: headers() });
    return res.data.data;
  },

  async skipSession(id) {
    const res = await axios.patch(`${API_URL}/planner/${id}/skip`, {}, { headers: headers() });
    return res.data.data;
  }
};

// ──── Weekly Targets ────
export const targetsApi = {
  async getTargets(weekDate) {
    const params = weekDate ? `?week=${weekDate}` : '';
    const res = await axios.get(`${API_URL}/weekly-targets${params}`, { headers: headers() });
    return res.data;
  },

  async getTarget(id) {
    const res = await axios.get(`${API_URL}/weekly-targets/${id}`, { headers: headers() });
    return res.data.data;
  },

  async createTarget(data) {
    const res = await axios.post(`${API_URL}/weekly-targets`, data, { headers: headers() });
    return res.data.data;
  },

  async updateTarget(id, data) {
    const res = await axios.put(`${API_URL}/weekly-targets/${id}`, data, { headers: headers() });
    return res.data.data;
  },

  async deleteTarget(id) {
    await axios.delete(`${API_URL}/weekly-targets/${id}`, { headers: headers() });
  },

  async createSubtask(targetId, data) {
    const res = await axios.post(`${API_URL}/weekly-targets/${targetId}/subtasks`, data, { headers: headers() });
    return res.data.data;
  },

  async updateSubtask(targetId, subtaskId, data) {
    const res = await axios.put(`${API_URL}/weekly-targets/${targetId}/subtasks/${subtaskId}`, data, { headers: headers() });
    return res.data.data;
  },

  async deleteSubtask(targetId, subtaskId) {
    await axios.delete(`${API_URL}/weekly-targets/${targetId}/subtasks/${subtaskId}`, { headers: headers() });
  },

  async toggleSubtask(targetId, subtaskId) {
    const res = await axios.patch(`${API_URL}/weekly-targets/${targetId}/subtasks/${subtaskId}/toggle`, {}, { headers: headers() });
    return res.data.data;
  }
};

// ──── Progress ────
export const progressApi = {
  async getProgress() {
    const res = await axios.get(`${API_URL}/progress`, { headers: headers() });
    return res.data.data;
  }
};
