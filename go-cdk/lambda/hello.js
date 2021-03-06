'use strict';

module.exports.handler = async (event) => {
    return {
        statusCode: 200,
        body: JSON.stringify(
            {
                message: 'Helo Serverless Lambda.',
                method: event.httpMethod,
                path: event.path,
                // input: event,
            },
            null,
            4
        ),
    };

    // Use this code if you don't use the http event with the LAMBDA-PROXY integration
    // return { message: 'Go Serverless v1.0! Your function executed successfully!', event };
};