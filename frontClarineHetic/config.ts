import * as dotenv from 'dotenv';
dotenv.config();

export const API_URL = process.env.API_URL || 'http://localhost:9070';
