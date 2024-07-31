"""
Utility to format the query result into the desired format (xml or opml)
"""

from typing import Any, Type
from xml.etree.ElementTree import Element, SubElement, tostring

from fastapi import HTTPException, Response, status

from src.schemas.lists import ListOut
from src.schemas.podcast import PodcastOut


def format_query_response(
    query_results: list[Type[Any]], search_format: str
) -> Response | HTTPException:
    """
    Return the podcasts in the given format or an HTTPException if format is not supported\
    Supported formats: JSON, OPML, XML
    """
    match search_format:
        case "opml":
            root = Element("opml", version="2.0")
            head = SubElement(root, "head")
            title = SubElement(head, "title")
            title.text = "Search Results"
            body = SubElement(root, "body")
            for element in query_results:
                body.append(element.to_opml())
            opml_content = tostring(root, encoding="utf-8").decode("utf-8")
            return Response(content=opml_content, media_type="text/xml")
        case "xml":
            if type(query_results[0]) == PodcastOut:
                root = Element("podcasts")
            elif type(query_results[0]) == ListOut:
                root = Element("podcasts lists")
            for element in query_results:
                root.append(element.to_xml())
            xml_content = tostring(root, encoding="utf-8").decode("utf-8")
            return Response(content=xml_content, media_type="text/xml")
        case "json":
            return query_results
        case _:
            raise HTTPException(
                detail="Format not supported",
                status_code=status.HTTP_406_NOT_ACCEPTABLE,
            )
