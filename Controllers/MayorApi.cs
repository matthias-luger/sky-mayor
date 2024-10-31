using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Http;
using Swashbuckle.AspNetCore.Annotations;
using Swashbuckle.AspNetCore.SwaggerGen;
using Newtonsoft.Json;
using Coflnet.Sky.Mayor.Attributes;
using Coflnet.Sky.Mayor.Models;

namespace Coflnet.Sky.Mayor.Controllers
{ 
    /// <summary>
    /// 
    /// </summary>
    [ApiController]
    public class MayorApiController : ControllerBase
    { 
        /// <summary>
        /// Get the the current mayor
        /// </summary>
        /// <remarks>Returns the name of the current mayor</remarks>
        /// <response code="200">OK</response>
        /// <response code="400">Bad Request</response>
        [HttpGet]
        [Route("/mayor/current")]
        [ValidateModelState]
        [SwaggerOperation("MayorCurrentGet")]
        [SwaggerResponse(statusCode: 200, type: typeof(ModelCandidate), description: "OK")]
        public virtual IActionResult MayorCurrentGet()
        {

            //TODO: Uncomment the next line to return response 200 or use other options such as return this.NotFound(), return this.BadRequest(..), ...
            // return StatusCode(200, default(ModelCandidate));
            //TODO: Uncomment the next line to return response 400 or use other options such as return this.NotFound(), return this.BadRequest(..), ...
            // return StatusCode(400);
            string exampleJson = null;
            exampleJson = "{\n  \"name\" : \"name\",\n  \"perks\" : [ {\n    \"name\" : \"name\",\n    \"description\" : \"description\"\n  }, {\n    \"name\" : \"name\",\n    \"description\" : \"description\"\n  } ],\n  \"key\" : \"key\"\n}";
            
            var example = exampleJson != null
            ? JsonConvert.DeserializeObject<ModelCandidate>(exampleJson)
            : default(ModelCandidate);
            //TODO: Change the data returned
            return new ObjectResult(example);
        }

        /// <summary>
        /// Get the name of the last mayor
        /// </summary>
        /// <remarks>Returns the name of the last mayor</remarks>
        /// <response code="200">OK</response>
        /// <response code="400">Bad Request</response>
        [HttpGet]
        [Route("/mayor/last")]
        [ValidateModelState]
        [SwaggerOperation("MayorLastGet")]
        [SwaggerResponse(statusCode: 200, type: typeof(string), description: "OK")]
        public virtual IActionResult MayorLastGet()
        {

            //TODO: Uncomment the next line to return response 200 or use other options such as return this.NotFound(), return this.BadRequest(..), ...
            // return StatusCode(200, default(string));
            //TODO: Uncomment the next line to return response 400 or use other options such as return this.NotFound(), return this.BadRequest(..), ...
            // return StatusCode(400);
            string exampleJson = null;
            
            var example = exampleJson != null
            ? JsonConvert.DeserializeObject<string>(exampleJson)
            : default(string);
            //TODO: Change the data returned
            return new ObjectResult(example);
        }

        /// <summary>
        /// Get names of all mayors
        /// </summary>
        /// <remarks>Returns all mayor names</remarks>
        /// <response code="200">OK</response>
        /// <response code="400">Bad Request</response>
        /// <response code="404">Not Found</response>
        [HttpGet]
        [Route("/mayor/names")]
        [ValidateModelState]
        [SwaggerOperation("MayorNamesGet")]
        [SwaggerResponse(statusCode: 200, type: typeof(List<string>), description: "OK")]
        public virtual IActionResult MayorNamesGet()
        {

            //TODO: Uncomment the next line to return response 200 or use other options such as return this.NotFound(), return this.BadRequest(..), ...
            // return StatusCode(200, default(List<string>));
            //TODO: Uncomment the next line to return response 400 or use other options such as return this.NotFound(), return this.BadRequest(..), ...
            // return StatusCode(400);
            //TODO: Uncomment the next line to return response 404 or use other options such as return this.NotFound(), return this.BadRequest(..), ...
            // return StatusCode(404);
            string exampleJson = null;
            exampleJson = "[ \"\", \"\" ]";
            
            var example = exampleJson != null
            ? JsonConvert.DeserializeObject<List<string>>(exampleJson)
            : default(List<string>);
            //TODO: Change the data returned
            return new ObjectResult(example);
        }

        /// <summary>
        /// Get the next mayor
        /// </summary>
        /// <remarks>Returns the mayor with the most votes in the current election. If there is currently no election, this returns null.</remarks>
        /// <response code="200">OK</response>
        /// <response code="400">Bad Request</response>
        /// <response code="404">Not Found</response>
        [HttpGet]
        [Route("/mayor/next")]
        [ValidateModelState]
        [SwaggerOperation("MayorNextGet")]
        [SwaggerResponse(statusCode: 200, type: typeof(ModelCandidate), description: "OK")]
        public virtual IActionResult MayorNextGet()
        {

            //TODO: Uncomment the next line to return response 200 or use other options such as return this.NotFound(), return this.BadRequest(..), ...
            // return StatusCode(200, default(ModelCandidate));
            //TODO: Uncomment the next line to return response 400 or use other options such as return this.NotFound(), return this.BadRequest(..), ...
            // return StatusCode(400);
            //TODO: Uncomment the next line to return response 404 or use other options such as return this.NotFound(), return this.BadRequest(..), ...
            // return StatusCode(404);
            string exampleJson = null;
            exampleJson = "{\n  \"name\" : \"name\",\n  \"perks\" : [ {\n    \"name\" : \"name\",\n    \"description\" : \"description\"\n  }, {\n    \"name\" : \"name\",\n    \"description\" : \"description\"\n  } ],\n  \"key\" : \"key\"\n}";
            
            var example = exampleJson != null
            ? JsonConvert.DeserializeObject<ModelCandidate>(exampleJson)
            : default(ModelCandidate);
            //TODO: Change the data returned
            return new ObjectResult(example);
        }
    }
}
