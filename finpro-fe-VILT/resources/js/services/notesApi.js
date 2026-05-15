import axios from 'axios';

// Gunakan URL API backend (Go)
const API_URL = 'http://localhost:8008/api/dashboard/notes';

// Ambil token dari localStorage atau cookie (sesuai setup aplikasi)
let authToken = null;

export const setAuthToken = (token) => {
  authToken = token;
};

function getHeaders() {
  return {
    Authorization: `Bearer ${authToken}`,
    'Content-Type': 'application/json'
  };
}

export const notesApi = {
  async getNotes() {
    const res = await axios.get(API_URL, { headers: getHeaders() });
    return res.data.data;
  },

  async searchNotes(query) {
    const res = await axios.get(`${API_URL}/search?q=${encodeURIComponent(query)}`, { headers: getHeaders() });
    return res.data.data;
  },

  async getNote(id) {
    const res = await axios.get(`${API_URL}/${id}`, { headers: getHeaders() });
    return res.data.data;
  },

  async getTemplates() {
    const res = await axios.get(`${API_URL}/templates`, { headers: getHeaders() });
    return res.data.data;
  },

  async createNote(data) {
    const res = await axios.post(API_URL, data, { headers: getHeaders() });
    return res.data.data;
  },

  async updateNote(id, data) {
    const res = await axios.put(`${API_URL}/${id}`, data, { headers: getHeaders() });
    return res.data.data;
  },

  async deleteNote(id) {
    const res = await axios.delete(`${API_URL}/${id}`, { headers: getHeaders() });
    return res.data;
  },

  async togglePin(id) {
    const res = await axios.put(`${API_URL}/${id}/pin`, {}, { headers: getHeaders() });
    return res.data.data;
  },

  async toggleArchive(id) {
    const res = await axios.put(`${API_URL}/${id}/archive`, {}, { headers: getHeaders() });
    return res.data.data;
  }
};
