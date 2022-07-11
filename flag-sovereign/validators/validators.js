const { body } = require('express-validator');

// validate POST body contains sdkkey and flags array
exports.validateFlagset = [body().isArray(), body("*.sdkKey").notEmpty(), body("*.flags").isArray()];