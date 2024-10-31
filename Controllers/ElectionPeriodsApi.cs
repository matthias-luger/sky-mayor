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
    public class ElectionPeriodsApiController : ControllerBase
    { 
        /// <summary>
        /// Inserts election periods
        /// </summary>
        /// <remarks>Endpoint to insert election periods, should only be used to insert missing/hisotical data</remarks>
        /// <param name="periods">the election periods that are going to be inserted</param>
        /// <response code="201">Created</response>
        /// <response code="400">Bad Request</response>
        [HttpPost]
        [Route("/electionPeriod")]
        [ValidateModelState]
        [SwaggerOperation("ElectionPeriodPost")]
        [SwaggerResponse(statusCode: 201, type: typeof(List<ModelElectionPeriod>), description: "Created")]
        public virtual IActionResult ElectionPeriodPost([FromBody]List<ModelElectionPeriod> periods)
        {

            //TODO: Uncomment the next line to return response 201 or use other options such as return this.NotFound(), return this.BadRequest(..), ...
            // return StatusCode(201, default(List<ModelElectionPeriod>));
            //TODO: Uncomment the next line to return response 400 or use other options such as return this.NotFound(), return this.BadRequest(..), ...
            // return StatusCode(400);
            string exampleJson = null;
            exampleJson = "[ {\n  \"candidates\" : [ {\n    \"name\" : \"name\",\n    \"perks\" : [ {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    }, {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    } ],\n    \"key\" : \"key\"\n  }, {\n    \"name\" : \"name\",\n    \"perks\" : [ {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    }, {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    } ],\n    \"key\" : \"key\"\n  } ],\n  \"winner\" : {\n    \"name\" : \"name\",\n    \"perks\" : [ {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    }, {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    } ],\n    \"key\" : \"key\"\n  },\n  \"year\" : 0,\n  \"start\" : \"start\",\n  \"end\" : \"end\",\n  \"id\" : \"id\"\n}, {\n  \"candidates\" : [ {\n    \"name\" : \"name\",\n    \"perks\" : [ {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    }, {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    } ],\n    \"key\" : \"key\"\n  }, {\n    \"name\" : \"name\",\n    \"perks\" : [ {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    }, {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    } ],\n    \"key\" : \"key\"\n  } ],\n  \"winner\" : {\n    \"name\" : \"name\",\n    \"perks\" : [ {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    }, {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    } ],\n    \"key\" : \"key\"\n  },\n  \"year\" : 0,\n  \"start\" : \"start\",\n  \"end\" : \"end\",\n  \"id\" : \"id\"\n} ]";
            
            var example = exampleJson != null
            ? JsonConvert.DeserializeObject<List<ModelElectionPeriod>>(exampleJson)
            : default(List<ModelElectionPeriod>);
            //TODO: Change the data returned
            return new ObjectResult(example);
        }

        /// <summary>
        /// Get election periods by timespan
        /// </summary>
        /// <remarks>Returns all election periods that took place in a given timespan</remarks>
        /// <param name="from">from The beginning of the selected timespan</param>
        /// <param name="to">The end of the selected timespan</param>
        /// <response code="200">OK</response>
        /// <response code="400">Bad Request</response>
        /// <response code="404">Not Found</response>
        [HttpGet]
        [Route("/electionPeriod/range")]
        [ValidateModelState]
        [SwaggerOperation("ElectionPeriodRangeGet")]
        [SwaggerResponse(statusCode: 200, type: typeof(List<ModelElectionPeriod>), description: "OK")]
        public virtual IActionResult ElectionPeriodRangeGet([FromQuery (Name = "from")][Required()]long from, [FromQuery (Name = "to")][Required()]long to)
        {

            //TODO: Uncomment the next line to return response 200 or use other options such as return this.NotFound(), return this.BadRequest(..), ...
            // return StatusCode(200, default(List<ModelElectionPeriod>));
            //TODO: Uncomment the next line to return response 400 or use other options such as return this.NotFound(), return this.BadRequest(..), ...
            // return StatusCode(400);
            //TODO: Uncomment the next line to return response 404 or use other options such as return this.NotFound(), return this.BadRequest(..), ...
            // return StatusCode(404);
            string exampleJson = null;
            exampleJson = "[ {\n  \"candidates\" : [ {\n    \"name\" : \"name\",\n    \"perks\" : [ {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    }, {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    } ],\n    \"key\" : \"key\"\n  }, {\n    \"name\" : \"name\",\n    \"perks\" : [ {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    }, {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    } ],\n    \"key\" : \"key\"\n  } ],\n  \"winner\" : {\n    \"name\" : \"name\",\n    \"perks\" : [ {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    }, {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    } ],\n    \"key\" : \"key\"\n  },\n  \"year\" : 0,\n  \"start\" : \"start\",\n  \"end\" : \"end\",\n  \"id\" : \"id\"\n}, {\n  \"candidates\" : [ {\n    \"name\" : \"name\",\n    \"perks\" : [ {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    }, {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    } ],\n    \"key\" : \"key\"\n  }, {\n    \"name\" : \"name\",\n    \"perks\" : [ {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    }, {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    } ],\n    \"key\" : \"key\"\n  } ],\n  \"winner\" : {\n    \"name\" : \"name\",\n    \"perks\" : [ {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    }, {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    } ],\n    \"key\" : \"key\"\n  },\n  \"year\" : 0,\n  \"start\" : \"start\",\n  \"end\" : \"end\",\n  \"id\" : \"id\"\n} ]";
            
            var example = exampleJson != null
            ? JsonConvert.DeserializeObject<List<ModelElectionPeriod>>(exampleJson)
            : default(List<ModelElectionPeriod>);
            //TODO: Change the data returned
            return new ObjectResult(example);
        }

        /// <summary>
        /// Get the election period of a certain year
        /// </summary>
        /// <remarks>Returns the election periods that took place in a given year</remarks>
        /// <param name="year">the searched year</param>
        /// <response code="200">OK</response>
        /// <response code="400">Bad Request</response>
        /// <response code="404">Not Found</response>
        [HttpGet]
        [Route("/electionPeriod/{year}")]
        [ValidateModelState]
        [SwaggerOperation("ElectionPeriodYearGet")]
        [SwaggerResponse(statusCode: 200, type: typeof(ModelElectionPeriod), description: "OK")]
        public virtual IActionResult ElectionPeriodYearGet([FromRoute (Name = "year")][Required]int year)
        {

            //TODO: Uncomment the next line to return response 200 or use other options such as return this.NotFound(), return this.BadRequest(..), ...
            // return StatusCode(200, default(ModelElectionPeriod));
            //TODO: Uncomment the next line to return response 400 or use other options such as return this.NotFound(), return this.BadRequest(..), ...
            // return StatusCode(400);
            //TODO: Uncomment the next line to return response 404 or use other options such as return this.NotFound(), return this.BadRequest(..), ...
            // return StatusCode(404);
            string exampleJson = null;
            exampleJson = "{\n  \"candidates\" : [ {\n    \"name\" : \"name\",\n    \"perks\" : [ {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    }, {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    } ],\n    \"key\" : \"key\"\n  }, {\n    \"name\" : \"name\",\n    \"perks\" : [ {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    }, {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    } ],\n    \"key\" : \"key\"\n  } ],\n  \"winner\" : {\n    \"name\" : \"name\",\n    \"perks\" : [ {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    }, {\n      \"name\" : \"name\",\n      \"description\" : \"description\"\n    } ],\n    \"key\" : \"key\"\n  },\n  \"year\" : 0,\n  \"start\" : \"start\",\n  \"end\" : \"end\",\n  \"id\" : \"id\"\n}";
            
            var example = exampleJson != null
            ? JsonConvert.DeserializeObject<ModelElectionPeriod>(exampleJson)
            : default(ModelElectionPeriod);
            //TODO: Change the data returned
            return new ObjectResult(example);
        }
    }
}
