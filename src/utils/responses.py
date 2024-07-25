"""
Utility to format the query result into the desired format (xml or opml)
"""

from xml.etree.ElementTree import Element, SubElement, tostring

from fastapi import HTTPException, Response, status

from src.models.podcast import Podcast


def format_podcasts_response(
    podcasts: list[Podcast], search_format: str
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
            title.text = "Podcast Search Results"
            body = SubElement(root, "body")
            for podcast in podcasts:
                body.append(podcast.to_opml())
            opml_content = tostring(root, encoding="utf-8").decode("utf-8")
            return Response(content=opml_content, media_type="text/xml")
        case "xml":
            root = Element("podcasts")
            for podcast in podcasts:
                root.append(podcast.to_xml())
            xml_content = tostring(root, encoding="utf-8").decode("utf-8")
            return Response(content=xml_content, media_type="text/xml")
        case "json":
            return podcasts
        case _:
            raise HTTPException(
                detail="Format not supported",
                status_code=status.HTTP_406_NOT_ACCEPTABLE,
            )
