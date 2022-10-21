import axios from 'axios';

const apiClient = {
  getFlags: async () => {
    let { data } = await axios.get('/api/flags');
    return data;
  },
  getFlag: async (id) => {
    let { data } = await axios.get(`/api/flags/${id}`);
    return data;
  },
  createFlag: async (newFlag) => {
    let { data } = await axios.post(`/api/flags`, newFlag);
    return data;
  },
  toggleFlag: async (id, newStatusObj) => {
    let { data } = await axios.patch(`/api/flags/${id}/toggle`, newStatusObj)
    return data;
  },
  editFlag: async (id, updatedFields) => {
    let { data } = await axios.patch(`/api/flags/${id}`, updatedFields);
    return data;
  },
  deleteFlag: async (id) => {
    let { data } = await axios.delete(`/api/flags/${id}`);
    return data;
  },
  getAudiences: async () => {
    let { data } = await axios.get('/api/audiences');
    return data;
  },
  getAudience: async(id) => {
    let { data } = await axios.get(`/api/audiences/${id}`);
    data.conditions = data.conditions.map(c => {
      const { Attribute, ...remainingFields } = c;
      return remainingFields;
    })

    return data;
  },
  createAudience: async (newAudience) => {
    let { data } = await axios.post('/api/audiences', newAudience);
    return data;
  },
  editAudience: async (id, updatedFields) => {
    let { data } = await axios.patch(`/api/audiences/${id}`, updatedFields);
    return data;
  },
  deleteAudience: async (id) => {
    let { data } = await axios.delete(`/api/audiences/${id}`);
    return data;
  },
  getAttributes: async () => {
    let { data } = await axios.get('/api/attributes');
    return data;
  },
  getAttribute: async (id) => {
    let { data } = await axios.get(`/api/attributes/${id}`);
    return data;
  },
  createAttribute: async (newAttr) => {
    let { data } = await axios.post(`/api/attributes`, newAttr)
    return data
  },
  deleteAttribute: async (id) => {
    let { data } = await axios.delete(`/api/attributes/${id}`)
    return data
  },
  getLogs: async () => {
    let { data } = await axios.get(`/api/auditlogs`);
    return data;
  },
  getSdkKey: async () => {
    let { data } = await axios.get(`/api/sdkkeys`);
    return data;
  },
  regenSdkKey: async (keyId) => {
    let { data } = await axios.delete(`/api/sdkkeys/${keyId}`);
    return data;
  }
}

export default apiClient;